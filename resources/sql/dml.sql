--SAVE ROLE
insert into accounts.roles(name, description) value ($1, $2);

--GET ROLE BY CODE
select r.name, r.description
from accounts.roles r
where r.name = $1;

select r.name, r.description
from accounts.roles r
order by r.name asc
limit $1 offset $2;

update accounts.roles
set name = $1,
    description = $2
where name = $3;
