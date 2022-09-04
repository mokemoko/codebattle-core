CREATE TABLE user
(
    id       VARCHAR(36)  NOT NULL,
    name     VARCHAR(100) NOT NULL,
    is_admin BOOL         NOT NULL DEFAULT FALSE,
    PRIMARY KEY (id)
);

CREATE TABLE contest
(
    id          VARCHAR(36)   NOT NULL,
    owner       VARCHAR(36)   NOT NULL,
    name        VARCHAR(100)  NOT NULL,
    description VARCHAR(1000),
    repository  VARCHAR(1000) NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (owner) REFERENCES user (id)
);

CREATE TABLE entry
(
    id         VARCHAR(36)   NOT NULL,
    user_id    VARCHAR(36)   NOT NULL,
    contest_id VARCHAR(36)   NOT NULL,
    name       VARCHAR(100)  NOT NULL,
    repository VARCHAR(1000) NOT NULL,
    status     int           NOT NULL,
    error      VARCHAR(1000),
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES user (id),
    FOREIGN KEY (contest_id) REFERENCES contest (id)
);

CREATE TABLE `match`
(
    id           VARCHAR(36) NOT NULL,
    entry_id     VARCHAR(36) NOT NULL,
    contest_id   VARCHAR(36) NOT NULL,
    `rank`       int         NOT NULL,
    before_score int         NOT NULL,
    after_score  int         NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (entry_id) REFERENCES entry (id),
    FOREIGN KEY (contest_id) REFERENCES contest (id)
);
