alter table university_reviews
    add column review_id text not null primary key;

alter table university_reviews
    drop column repliesCount;

alter table university_reviews
    add parent_review_id text references university_reviews (review_id);