DROP TABLE IF EXISTS company
;
CREATE TABLE company(
    company_cd varchar(5) NOT NULL,
    company_name varchar(256) NOT NULL,
    founded_year integer NOT NULL,
    status varchar(1) NOT NULL default 0,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    revision integer NOT NULL,
    CONSTRAINT company_pkc PRIMARY KEY(company_cd)
)
;
