-- name: GetOpenDays :many

select u.name, od.description, od.address, od.link, od.date
from university_open_days od
         join universities u on u.id = od.university_id
where u.id = $1;
