CREATE OR REPLACE FUNCTION orders_get()
	RETURNS json
	LANGUAGE plpgsql
AS $function$
DECLARE
	_response JSONB;
BEGIN
	SELECT
		COALESCE(og.s, '[]')
	FROM
	(
		SELECT json_agg(ag.*)::JSONB s
		FROM (
			SELECT o.id, o.uid_order, o.created, o.status, r.member_fixed
			FROM orders o, rooms r
			WHERE o.uid_order = r.uid_room
		) ag
	) og
	INTO _response;

	RETURN _response;
END;
$function$