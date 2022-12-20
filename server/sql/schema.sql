DROP TABLE IF EXISTS user;
CREATE TABLE user
(
    id         VARCHAR(36)  NOT NULL,
    name       VARCHAR(100) NOT NULL,
    icon       VARCHAR(100) NOT NULL,
    is_admin   BOOL         NOT NULL DEFAULT FALSE,
    token      VARCHAR(100),
    created_at TIMESTAMP    NOT NULL,
    updated_at TIMESTAMP    NOT NULL,
    PRIMARY KEY (id)
);

DROP TABLE IF EXISTS contest;
CREATE TABLE contest
(
    id          VARCHAR(36)   NOT NULL,
    owner       VARCHAR(36)   NOT NULL,
    name        VARCHAR(100)  NOT NULL,
    description VARCHAR(1000),
    repository  VARCHAR(1000) NOT NULL,
    created_at  TIMESTAMP     NOT NULL,
    updated_at  TIMESTAMP     NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (owner) REFERENCES user (id)
);

DROP TABLE IF EXISTS entry;
CREATE TABLE entry
(
    id         VARCHAR(36)   NOT NULL,
    user_id    VARCHAR(36)   NOT NULL,
    contest_id VARCHAR(36)   NOT NULL,
    name       VARCHAR(100)  NOT NULL,
    repository VARCHAR(1000) NOT NULL,
    status     INT           NOT NULL,
    error      VARCHAR(1000),
    score      INT           NOT NULL,
    created_at TIMESTAMP     NOT NULL,
    updated_at TIMESTAMP     NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES user (id),
    FOREIGN KEY (contest_id) REFERENCES contest (id)
);

# TODO: 正規化
DROP TABLE IF EXISTS `match`;
CREATE TABLE `match`
(
    id           VARCHAR(36) NOT NULL,
    entry_id     VARCHAR(36) NOT NULL,
    contest_id   VARCHAR(36) NOT NULL,
    status       INT         NOT NULL DEFAULT 0,
    `rank`       INT         NOT NULL,
    before_score INT         NOT NULL,
    after_score  INT         NOT NULL,
    created_at   TIMESTAMP   NOT NULL,
    updated_at   TIMESTAMP   NOT NULL,
    PRIMARY KEY (id, entry_id),
    FOREIGN KEY (entry_id) REFERENCES entry (id),
    FOREIGN KEY (contest_id) REFERENCES contest (id)
);
