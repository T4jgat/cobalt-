CREATE TABLE rentals (
                         id SERIAL PRIMARY KEY,
                         user_id INT NOT NULL,
                         car_id INT NOT NULL,
                         start_date DATE NOT NULL,
                         end_date DATE NOT NULL,
                         status VARCHAR(50) NOT NULL,
                         created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);