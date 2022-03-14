CREATE TABLE IF NOT EXISTS laptop_loan (
    id BLOB,
    student_id TEXT NOT NULL,
    start_date DATETIME NOT NULL,
    end_date DATETIME NOT NULL,
    file BLOB,
    PRIMARY KEY(id)
);