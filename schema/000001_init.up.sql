CREATE TABLE storage_links_tab
(
    id            serial       not null unique,
    short_link    varchar(255) not null unique,
    original_link text
);