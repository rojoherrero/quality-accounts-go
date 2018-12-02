--SAVE ROLE
insert into accounts.roles(code, description) value ($1, $2);

--GET ROLE BY CODE
select r.code, r.description
from accounts.roles r
where r.code = $1;

select r.code, r.description
from accounts.roles r
order by r.code asc
limit $1 offset $2;

update accounts.roles
set code = $1,
    description = $2
where code = $3;
