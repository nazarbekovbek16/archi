CREATE TABLE IF NOT EXISTS history_books (
    id bigserial PRIMARY KEY,
    user_id int,
    book_id int,
    duration int,
    is_paid bool,
    is_given bool,
    price int
);