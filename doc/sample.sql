CREATE TABLE IF NOT EXISTS services (
	id SERIAL PRIMARY KEY,
	name VARCHAR(50),
    cron VARCHAR(50) NOT NULL,
	url VARCHAR(50) NOT NULL,
	type VARCHAR(15) NOT NULL
);

CREATE TABLE IF NOT EXISTS results (
	service_id INT,
    success BOOLEAN,
    reason TEXT,
    cron_time TIMESTAMP NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT current_timestamp,
    FOREIGN KEY (service_id) REFERENCES services(id),
    PRIMARY KEY (service_id, cron_time)
);