import mysql.connector as mysql
from faker import Faker
from time import time
import numpy as np
import base64

fake = Faker()


class KeyGen:
    min: int = 0
    max: int = 100
    size: int = 800

    # We are simulating a Kyber public key by creating an array of 800 uint32 values
    def gen_public_key(self) -> str:
        array = np.random.randint(self.min, self.max, size=self.size, dtype=np.uint32)
        byte_data = array.tobytes()
        base64_string = base64.b64encode(byte_data).decode("utf-8")
        return base64_string


class Benchmark:
    iterations: int = 100000
    batch_size: int = 10000
    duration: float = 0.0

    def __init__(self, add_fn, read_fn, iter_buf):
        self.add_fn = add_fn
        self.read_fn = read_fn
        self.iter_buf = iter_buf

    def measure(self, func):
        start = time()
        func()
        end = time()
        self.duration += (end - start)

    def iterate(self):
        for i in range(self.iterations):
            self.measure(self.add_fn)

            if i >= self.batch_size and i % self.batch_size == 0:
                # Find the average time taken to add a user
                start = time()
                self.read_fn()
                end = time()
                read_avg = (end - start) * 1000
                print(f"Average read time per {i} iterations: {read_avg:.4f} ms")

                # Find the average time taken to create a user
                avg_add = (self.duration / i) * 1000
                print(f"Average add time per {i} iterations: {avg_add:.4f} ms")


class Database:
    user = "user"
    password = "password"
    database = "mysql"
    host = "localhost"
    port = 3306
    table_name: str = "Benchmark"
    columns: str = f"""
    id INT AUTO_INCREMENT PRIMARY KEY,
    publickey TEXT,
    username VARCHAR(255),
    email VARCHAR(255)"""
    user_count: int = 0
    current_user: str

    def __init__(self):
        start_time = time()
        wait_warned = False

        print("Waiting for MySQL connection...")
        while True:
            try:
                self.connect()
                if self.conn:
                    print("Connected to database")
                    break
            except Exception:
                if time() - start_time >= 5 and not wait_warned:
                    print("Still waiting for MySQL connection...")
                    wait_warned = True
                elif time() - start_time >= 10:
                    print("Connection failed")
                    exit()


    def connect(self):
        self.conn = mysql.connect(
            user=self.user,
            password=self.password,
            database=self.database,
            host=self.host,
            port=self.port,
        )
        self.cursor = self.conn.cursor()

    def create_table(self):
        self.cursor.execute(
            f"""CREATE TABLE
            IF NOT EXISTS
            {self.table_name}({self.columns})"""
        )

    def create_user(self):
        email = fake.email()
        username = fake.user_name()
        public_key = KeyGen().gen_public_key()

        self.user_count += 1
        self.current_user = username

        self.cursor.execute(
            f"""INSERT INTO Benchmark
            VALUES (
            NULL,
            '{public_key}',
            '{username}',
            '{email}')"""
        )

    def get_users(self):
        self.cursor.execute(f"""SELECT * FROM {self.table_name}""")
        for row in self.cursor.fetchall():
            print(row)

    def get_current_user(self):
        self.cursor.execute(
            f"""SELECT * FROM {self.table_name}
            WHERE username = '{self.current_user}'"""
        )

        # To avoid result not used error
        self.cursor.fetchall()
