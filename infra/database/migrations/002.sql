CREATE TRIGGER populate_weighted_doc_trigger_func_temp
AFTER
INSERT on youtube_data for EACH ROW EXECUTE FUNCTION populate_weighted_doc_func_temp();


DROP FUNCTION populate_weighted_doc_func() CASCADE;
CREATE OR REPLACE PROCEDURE populate_weighted_doc_func() RETURNS trigger AS $$ begin
update youtube_data
set search_weighted_doc = setweight(to_tsvector(title), 'A') || setweight(to_tsvector(description), 'B')
where search_weighted_doc is null;
RETURN;
END $$ LANGUAGE plpgsql;
select populate_weighted_doc_func_temp();
select plainto_tsquery('surfing | wave | pizzo');
select *
from youtube_data
where to_tsquery('pizzo | wave') @@ search_doc;
select *
from youtube_data
where to_tsquery('summer | beginner') @@ search_weighted_doc
order by ts_rank(
    search_weighted_doc,
    to_tsquery('summer & beginner')
  ) desc;
select setweight(to_tsvector(title), 'A') || setweight(to_tsvector(description), 'B')
from youtube_data;
select *
from youtube_data;
---- to fix immutaable error
alter table youtube_data
add column search_weighted_doc tsvector;
alter table youtube_data
add column search_doc tsvector;
update youtube_data
set search_weighted_doc = setweight(to_tsvector(title), 'A') || setweight(to_tsvector(description), 'B');
update youtube_data
set search_doc = to_tsvector(title) || to_tsvector(description);
CREATE TRIGGER populate_weighted_doc_trigger_func_temp
AFTER
INSERT on youtube_data for EACH ROW EXECUTE FUNCTION populate_weighted_doc_func_temp();
DROP FUNCTION populate_weighted_doc_func() CASCADE
CREATE OR REPLACE PROCEDURE populate_weighted_doc_func() RETURNS trigger AS $$ begin
update youtube_data
set search_weighted_doc = setweight(to_tsvector(title), 'A') || setweight(to_tsvector(description), 'B')
where search_weighted_doc is null;
RETURN;
END $$ LANGUAGE plpgsql;
select populate_weighted_doc_func_temp() create index search_weights_idx on youtube_data using GIN(search_doc);
create index search_weights__idx on youtube_data using GIN(search_weighted_doc);
select *
from youtube_data
where to_tsquery('pizzo | wave') @@ search_doc;
select *
from youtube_data
where to_tsquery('summer | beginner') @@ search_weighted_doc
order by ts_rank(
    search_weighted_doc,
    to_tsquery('summer & beginner')
  ) desc;
select setweight(to_tsvector(title), 'A') || setweight(to_tsvector(description), 'B')
from youtube_data;
select title,
  ts_rank(
    to_tsvector(title || description),
    to_tsquery('java | sale | in')
  ) ranks,
  description
from courses
order by ranks desc;
select *
from courses
where to_tsquery('java | sale') @@ to_tsvector(title || description)
select courses.title,
  courses.description,
  setweight(ts_rank(to_tsvector(courses.title), query), 1) as rank_title,
  setweight(
    ts_rank(to_tsvector(courses.description), query),
    2
  ) as rank_description
from courses,
  to_tsvector(courses.title || '' || courses.description) document,
  to_tsquery('java | sale') query
where query @@ document
order by rank_title,
  rank_description DESC;