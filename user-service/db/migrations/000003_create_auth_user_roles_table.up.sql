CREATE TABLE auth_user_roles (
                                 user_id UUID REFERENCES auth_users(id) ON DELETE CASCADE,
                                 role_id INT REFERENCES auth_roles(id) ON DELETE CASCADE,
                                 PRIMARY KEY (user_id, role_id)
);
