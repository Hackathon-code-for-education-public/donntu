-- name: GetOpenDays :many

select u.name, od.description, od.address, od.link, od.date
from university_open_days od
         join universities u on u.id = od.university_id
where u.id = $1;

-- name: GetReviews :many
select *
from university_reviews ur
where ur.university_id = $3
offset $1 limit $2;

-- name: AddPanorama :one
insert into university_panoramas (university_id, address, name, firstlocation, secondlocation, type)
values ($1, $2, $3, $4, $5, $6) returning *;

-- name: GetPanoramas :many
select * from university_panoramas p where university_id = $1 and type = $2 order by p.name;