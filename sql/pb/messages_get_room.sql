CREATE OR REPLACE FUNCTION messages_get_room(_uidr character varying)
	RETURNS json
	LANGUAGE plpgsql
AS $function$
DECLARE
	_m messages;
	_response JSONB;
BEGIN
	SELECT *
	FROM messages
	WHERE uid_room = _uidr
	INTO _m;

	IF _m.id ISNULL THEN
		RAISE EXCEPTION 'сообщений с таким uid группы не существует';
	END IF;

	SELECT
		COALESCE(mgr.s, '[]')
	FROM
	(
		SELECT json_agg(ag.*)::JSONB s
		FROM (
			SELECT m.id, m.uid_room, m.uid_user, m.message, m.created
			FROM messages m
			WHERE m.uid_room = _uidr
		) ag
	) mgr
	INTO _response;

	RETURN _response;
END;
$function$