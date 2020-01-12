select Date,
    case Venue
when 'H' then 'Home' else 'Away' end as Venue,
    Opposition,
    case Points
when 3 then 'Win'
when 0 then 'Loss'
else 'Draw' end as Result,
    GoalsFor,
    GoalsAgainst
from points
limit 5;