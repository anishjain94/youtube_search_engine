select to_tsvector(title || description) from courses


select title, ts_rank(to_tsvector(title || description), to_tsquery('java | sale | in')) ranks, description from courses order by ranks desc


select * from courses where to_tsquery('java | sale') 
@@ to_tsvector(title || description)


select 
	courses.title, 
	courses.description, 
	ts_rank(to_tsvector(courses.title), query) as rank_title,
	ts_rank(to_tsvector(courses.description), query) as rank_description
from 
	courses,
	to_tsvector(courses.title || '' || courses.description) document,
	to_tsquery('java | sale') query
where query @@ document
order by rank_title, rank_description DESC


CREATE EXTENSION pg_trgm


SELECT 
    SIMILARITY('ani jain', 'anish kumar jain') as similarity


create index course_search_indx on courses using GIN (to_tsvector(courses.title  || courses.description ));




select title || ' ' || description from courses

SELECT plainto_tsquery('java in a nutshell');


SELECT to_tsquery('java | sale | in');


select to_tsvector(description) from courses