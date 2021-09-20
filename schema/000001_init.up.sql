CREATE TABLE storage_links_tab
(
    short_link    varchar(255) not null unique,
    original_link text
);