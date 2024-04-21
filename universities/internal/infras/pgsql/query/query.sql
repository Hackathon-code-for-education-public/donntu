-- name: GetOpenDays :many

select u.name, od.description, od.address, od.link, od.date
from university_open_days od
         join universities u on u.id = od.university_id
where u.id = $1;

-- name: CreateReview :one
insert into university_reviews(university_id,
                               author_status,
                               sentiment, date,
                               text, review_id, parent_review_id)
values ($1, $2, $3, $4, $5, $6, $7)
returning *;

-- name: GetReviews :many
select r.*,
       (select count(*)
        from university_reviews
        where parent_review_id = r.review_id) as reply_count
from university_reviews r
where r.university_id = $3
  and r.parent_review_id is null
group by r.review_id, r.date
order by r.date
offset $1 limit $2;

-- name: GetReviewsByParent :many
select *
from university_reviews r
where r.parent_review_id = $1
order by r.date;

-- name: GetReview :one
select r.*, (select count(*) from university_reviews where parent_review_id = r.review_id) as reply_count
from university_reviews r
where r.review_id = $1
limit 1;

-- name: AddPanorama :one
insert into university_panoramas (university_id, address, name, firstlocation, secondlocation, type)
values ($1, $2, $3, $4, $5, $6)
returning *;

-- name: GetPanoramas :many
select *
from university_panoramas p
where university_id = $1
  and type = $2
order by p.name;

-- name: SearchUniversities :many
select *
from universities
where name ilike $1
   or long_name ilike $2
order by name;

-- name: GetTopOfUniversities :many
select *
from universities u
order by u.rating desc
limit $1;

-- name: GetUniversities :many
select *
from universities u
order by u.name
offset $1 limit $2;

-- name: GetUniversity :one
select *
from universities
where id = $1
limit 1;
