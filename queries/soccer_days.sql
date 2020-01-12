select date_part('dow', Date) as DayOfWeek,
count(date_part('dow', Date)) as Days
from points
group by date_part('dow', Date)
order by count(date_part('dow', Date)) desc;