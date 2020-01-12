declare @start as date;
declare @finish as date;
declare @team as varchar(50);

set @start = '%s';
set @finish = '%s';
set @team = '%s';

select * from points
where team = @team
and Date between @start and @finish;