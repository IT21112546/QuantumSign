# Import the Libraries
import numpy as np
import pandas as pd
from flask import Flask, request, jsonify
from sklearn.model_selection import train_test_split
from sklearn.linear_model import LogisticRegression
from sklearn.tree import DecisionTreeClassifier
from sklearn.svm import SVC
from sklearn.ensemble import RandomForestClassifier
from sklearn.preprocessing import StandardScaler, OneHotEncoder
from sklearn.compose import ColumnTransformer
from sklearn.metrics import roc_auc_score
from sklearn.pipeline import Pipeline
import ipaddress
import joblib

# Initialize Flask app
app = Flask(__name__)

# Load the Dataset
data = pd.read_csv('/kaggle/input/rba-dataset/rba-dataset.csv', nrows=1000000)

# Data Preprocessing/Cleaning
data['Login Hour'] = pd.to_datetime(data['Login Timestamp']).dt.hour
data['Is Account Takeover'] = data['Is Account Takeover'].astype(np.uint8)
data['Is Attack IP'] = data['Is Attack IP'].astype(np.uint8)
data['Login Successful'] = data['Login Successful'].astype(np.uint8)
data = data.drop(columns=["Round-Trip Time [ms]", 'Region', 'City', 'Login Timestamp', 'index'])

# Convert Strings to Integers
data['User Agent String'], _ = pd.factorize(data['User Agent String'])
data['Browser Name and Version'], _ = pd.factorize(data['Browser Name and Version'])
data['OS Name and Version'], _ = pd.factorize(data['OS Name and Version'])

# Convert IP Addresses to Integers
def ip_to_int(ip):
    return int(ipaddress.ip_address(ip))

data['IP Address'] = data['IP Address'].apply(ip_to_int)

# Encoding Categorical & Numerical Variables
categorical_cols = ['Country', 'Device Type']
numeric_cols = ['ASN', 'Login Hour', 'IP Address', 'User Agent String', 'Browser Name and Version', 'OS Name and Version']

# Splitting the Dataset Into Train/Test
features = data.drop(['Is Attack IP', 'Is Account Takeover'], axis=1)
labels = data['Is Account Takeover']

X_train, X_test, y_train, y_test = train_test_split(features, labels, test_size=0.2, random_state=42)

# Preprocessors
preprocessor = ColumnTransformer(
    transformers=[
        ('num', StandardScaler(), numeric_cols),
        ('cat', OneHotEncoder(), categorical_cols)
    ])

# Classifiers
classifiers = {
    'logistic_regression': LogisticRegression(max_iter=1000),
    'decision_tree': DecisionTreeClassifier(),
    'svm': SVC(probability=True),
    'random_forest': RandomForestClassifier()
}

# A function to choose classifiers
def make_pipeline(classifier_key):
    if classifier_key in classifiers:
        clf = Pipeline(steps=[
            ('preprocessor', preprocessor),
            ('classifier', classifiers[classifier_key])
        ])
        return clf
    else:
        raise ValueError(f"Classifier {classifier_key} is not defined")

# Train Random Forest model (or any classifier you prefer)
classifier_key = 'random_forest'
pipeline = make_pipeline(classifier_key)
pipeline.fit(X_train, y_train)

# Save the trained model to a file
joblib.dump(pipeline, 'model.pkl')

# Load the trained model
model = joblib.load('model.pkl')

def predict():
    # Get JSON data from the HTTP request
    input_data = request.get_json()
    input_df = pd.DataFrame([input_data])

    # Convert IP address to integer
    if 'IP Address' in input_df.columns:
        input_df['IP Address'] = input_df['IP Address'].apply(ip_to_int)

    # Preprocessing: Converting factors for categorical columns
    if 'User Agent String' in input_df.columns:
        input_df['User Agent String'], _ = pd.factorize(input_df['User Agent String'])
    if 'Browser Name and Version' in input_df.columns:
        input_df['Browser Name and Version'], _ = pd.factorize(input_df['Browser Name and Version'])
    if 'OS Name and Version' in input_df.columns:
        input_df['OS Name and Version'], _ = pd.factorize(input_df['OS Name and Version'])

        prediction = model.predict(input_df)
        prediction_prob = model.predict_proba(input_df)[:, 1]

        # Return the result as a JSON response
        return jsonify({
            'prediction': int(prediction[0]),
            'probability': float(prediction_prob[0])
        })
