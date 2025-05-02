CREATE TABLE reviews (
                         id         UUID PRIMARY KEY,
                         salon_id   UUID NOT NULL,
                         user_id    UUID NOT NULL,
                         content    TEXT NOT NULL,
                         rating     INT  CHECK (rating BETWEEN 1 AND 5),
                         created_at TIMESTAMP NOT NULL
);
