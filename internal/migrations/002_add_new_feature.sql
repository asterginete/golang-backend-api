-- Create the reviews table
CREATE TABLE reviews (
    id SERIAL PRIMARY KEY,
    book_id INT NOT NULL REFERENCES books(id),
    user_id INT NOT NULL REFERENCES users(id),
    rating INT CHECK (rating >= 1 AND rating <= 5), -- Assuming a rating system from 1 to 5
    comment TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Add an average_rating column to the books table to store the average rating for each book
ALTER TABLE books ADD COLUMN average_rating FLOAT DEFAULT 0;
