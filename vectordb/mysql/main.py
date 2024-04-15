from utils import Database, Benchmark

# Initialize the database
db = Database()


def main():

    # Prepare the table
    db.create_table()

    # Start benchmarking
    benchmark = Benchmark(db.create_user, db.get_current_user, db.user_count)
    benchmark.iterate()


if __name__ == "__main__":
    main()
