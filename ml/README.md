+ A runnable version of this notebook is available on [Kaggle](https://www.kaggle.com/code/kottunaana/authentication-anomaly-detection-model).

# Problem Statement
+ In the field of authentication anomaly detection, there has been two very good papers that we will be basing our research on:
    1. [Who Are You? By Freeman et al.](https://iris.unica.it/retrieve/handle/11584/133076/390585/freeman16-ndss-final.pdf)
    2. [Risk Based Authentication for OpenStack By Vincent et al.](https://arxiv.org/pdf/2303.12361)

### Freeman et al.
+ Freeman et al. were the first in this field of research to discuss and come up with some kind of a way to identify suspicious login attempts.
+ Since Freeman et al. were working on LinkedIn's authentication dataset, which was private, we were never given the chance to exactly reproduce their methodology.
+ Also, Freeman et al. did not use a supervised learning approach to detect anomalies, they were using their own custom methodology using statistical algorithms.
+ They also claimed to have produced a statistical model that detects anomalies with a 96% accuracy.
+ The statistical model however was not using some sort of a Bayesian approach, their main methodology revolved around calculating a risk score using a static statistical equation which they derived in their paper.
+ Therefore, there is no 'learning' in their methodology, it is a static model that calculates a risk score for each login attempt.
+ The risk score is then compared to a threshold to determine if the login attempt is suspicious or not.

### Vincent et al.
+ Then comes Vincent et al. who attempted to reproduce Freeman et al.'s methodology using a [public dataset](https://www.kaggle.com/datasets/dasgroup/rba-dataset/data?select=rba-dataset.csv).
+ Their research objectives were to:
    1. Reproduce Freeman et al.'s methodology using this new dataset.
    2. Improve [OpenStack's](https://www.openstack.org/) authentication system using Freeman et al.'s methodology.
+ Their methodology also calculated a risk score for each login attempt using the same statistical equation as Freeman et al.

# Research Contribution
+ We will attempt to improve both Freeman et al.'s and Vincent et al.'s methodologies by using a supervised learning approach.
+ We will also attempt to develop a model with an increased accuracy compared to Freeman et al.'s 96% accuracy.

# Methodology
+ Our methodology is to use the following supervised learning classification algorithms:
    1. Logistic Regression
    2. Random Forest
    3. Support Vector Machine
    4. Decision Tree

+ We will compare the Area Under the Curve (AUC) of the Receiver Operating Characteristic (ROC) curve of each of these algorithms to determine which one is the best.

+ Freeman et al. and Vincent et al. both were relying on the 'City' and 'Region' features, however, we believe that these features are not very useful in determining the authenticity of a login attempt.
+ Instead, we believe that the addition of a 'Login Hour' (the hour of the day out of 24 hours) will provide more context during anomaly classification.
+ This is because it is possible for attackers to wait until the middle of the night to launch an attack, and this feature will help us identify such attacks.
+ Moreover, Freeman et al. had proposed this feature in their paper, however, Vincent et al did not seem to have used it in their research.
+ The 'Login Hour' feature will be extracted from the 'Time' feature, which is a timestamp of the login attempt.

# Research Findings
+ We trained our models using the said methodology, and these were the Area Under the Curve (AUC) of the Receiver Operating Characteristic (ROC) curve for each of the algorithms:

| Model                     | AUC                   |
|---------------------------|-----------------------|
| Logistic Regression       | 0.916850419422958     |
| Random Forest             | 0.8332983328083254    |
| Support Vector Machine    | 0.9879314856389512    |
| Decision Tree             | 0.8333308332958327    |

+ We can see that the Support Vector Machine model has the highest AUC with an accuracy of 98.79%.
+ This is greater than Freeman et al.'s 96% accuracy, and therefore, we can conclude that our model is better than Freeman et al.'s model.
