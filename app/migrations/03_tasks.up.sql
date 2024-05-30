CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS tasks (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  user_id INT REFERENCES users(id) ON DELETE CASCADE,
  category_id UUID REFERENCES categories(id) ON DELETE CASCADE,
  title VARCHAR(255) NOT NULL,
  description TEXT DEFAULT '',
  priority INT DEFAULT 1 CHECK (priority IN (1, 2, 3)), -- 1-basic, 2-important, 3-extremely important
  status INT DEFAULT 1 CHECK (status IN (1, 2, 3)),     -- 1-to_do, 2-progress, 3-done
  deadline TIMESTAMP NOT NULL CHECK (deadline > CURRENT_TIMESTAMP),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
