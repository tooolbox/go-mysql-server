package analyzer

import (
	"github.com/liquidata-inc/go-mysql-server/sql"
	"github.com/liquidata-inc/go-mysql-server/sql/plan"
	"strings"
)

var validColumnDefaultFuncs = map[string]struct{}{
	"abs": {},
	"acos": {},
	"adddate": {},
	"addtime": {},
	"aes_decrypt": {},
	"aes_encrypt": {},
	"any_value": {},
	"ascii": {},
	"asin": {},
	"atan2": {},
	"atan": {},
	"avg": {},
	"benchmark": {},
	"bin": {},
	"bin_to_uuid": {},
	"bit_and": {},
	"bit_length": {},
	"bit_count": {},
	"bit_or": {},
	"bit_xor": {},
	"can_access_column": {},
	"can_access_database": {},
	"can_access_table": {},
	"can_access_view": {},
	"cast": {},
	"ceil": {},
	"ceiling": {},
	"char": {},
	"char_length": {},
	"character_length": {},
	"charset": {},
	"coalesce": {},
	"coercibility": {},
	"collation": {},
	"compress": {},
	"concat": {},
	"concat_ws": {},
	"connection_id": {},
	"conv": {},
	"convert": {},
	"convert_tz": {},
	"cos": {},
	"cot": {},
	"count": {},
	"crc32": {},
	"cume_dist": {},
	"curdate": {},
	"current_role": {},
	"current_timestamp": {},
	"curtime": {},
	"database": {},
	"date": {},
	"date_add": {},
	"date_format": {},
	"date_sub": {},
	"datediff": {},
	"day": {},
	"dayname": {},
	"dayofmonth": {},
	"dayofweek": {},
	"dayofyear": {},
	"default": {},
	"degrees": {},
	"dense_rank": {},
	"elt": {},
	"exp": {},
	"export_set": {},
	"extract": {},
	"extractvalue": {},
	"field": {},
	"find_in_set": {},
	"first_value": {},
	"floor": {},
	"format": {},
	"format_bytes": {},
	"format_pico_time": {},
	"found_rows": {},
	"from_base64": {},
	"from_days": {},
	"from_unixtime": {},
	"geomcollection": {},
	"geometrycollection": {},
	"get_dd_column_privileges": {},
	"get_dd_create_options": {},
	"get_dd_index_sub_part_length": {},
	"get_format": {},
	"get_lock": {},
	"greatest": {},
	"group_concat": {},
	"grouping": {},
	"gtid_subset": {},
	"gtid_subtract": {},
	"hex": {},
	"hour": {},
	"icu_version": {},
	"if": {},
	"ifnull": {},
	"in": {},
	"inet_aton": {},
	"inet_ntoa": {},
	"inet6_aton": {},
	"inet6_ntoa": {},
	"insert": {},
	"instr": {},
	"internal_auto_increment": {},
	"internal_avg_row_length": {},
	"internal_check_time": {},
	"internal_checksum": {},
	"internal_data_free": {},
	"internal_data_length": {},
	"internal_dd_char_length": {},
	"internal_get_comment_or_error": {},
	"internal_get_enabled_role_json": {},
	"internal_get_hostname": {},
	"internal_get_username": {},
	"internal_get_view_warning_or_error": {},
	"internal_index_column_cardinality": {},
	"internal_index_length": {},
	"internal_is_enabled_role": {},
	"internal_is_mandatory_role": {},
	"internal_keys_disabled": {},
	"internal_max_data_length": {},
	"internal_table_rows": {},
	"internal_update_time": {},
	"interval": {},
	"is_free_lock": {},
	"is_ipv4": {},
	"is_ipv4_compat": {},
	"is_ipv4_mapped": {},
	"is_ipv6": {},
	"is_used_lock": {},
	"is_uuid": {},
	"isnull": {},
	"json_array": {},
	"json_array_append": {},
	"json_array_insert": {},
	"json_arrayagg": {},
	"json_contains": {},
	"json_contains_path": {},
	"json_depth": {},
	"json_extract": {},
	"json_insert": {},
	"json_keys": {},
	"json_length": {},
	"json_merge": {},
	"json_merge_patch": {},
	"json_merge_preserve": {},
	"json_object": {},
	"json_objectagg": {},
	"json_overlaps": {},
	"json_pretty": {},
	"json_quote": {},
	"json_remove": {},
	"json_replace": {},
	"json_schema_valid": {},
	"json_schema_validation_report": {},
	"json_search": {},
	"json_set": {},
	"json_storage_free": {},
	"json_storage_size": {},
	"json_table": {},
	"json_type": {},
	"json_unquote": {},
	"json_valid": {},
	"json_value": {},
	"lag": {},
	"last_insert_id": {},
	"last_value": {},
	"lcase": {},
	"lead": {},
	"least": {},
	"left": {},
	"length": {},
	"linestring": {},
	"ln": {},
	"load_file": {},
	"localtimestamp": {},
	"locate": {},
	"log": {},
	"log10": {},
	"log2": {},
	"lower": {},
	"lpad": {},
	"ltrim": {},
	"make_set": {},
	"makedate": {},
	"maketime": {},
	"master_pos_wait": {},
	"max": {},
	"mbrcontains": {},
	"mbrcoveredby": {},
	"mbrcovers": {},
	"mbrdisjoint": {},
	"mbrequals": {},
	"mbrintersects": {},
	"mbroverlaps": {},
	"mbrtouches": {},
	"mbrwithin": {},
	"md5": {},
	"microsecond": {},
	"mid": {},
	"min": {},
	"minute": {},
	"mod": {},
	"month": {},
	"monthname": {},
	"multilinestring": {},
	"multipoint": {},
	"multipolygon": {},
	"name_const": {},
	"now": {},
	"nth_value": {},
	"ntile": {},
	"nullif": {},
	"oct": {},
	"octet_length": {},
	"ord": {},
	"percent_rank": {},
	"period_add": {},
	"period_diff": {},
	"pi": {},
	"point": {},
	"polygon": {},
	"position": {},
	"pow": {},
	"power": {},
	"ps_current_thread_id": {},
	"ps_thread_id": {},
	"quarter": {},
	"quote": {},
	"radians": {},
	"rand": {},
	"random_bytes": {},
	"rank": {},
	"regexp_instr": {},
	"regexp_like": {},
	"regexp_replace": {},
	"regexp_substr": {},
	"release_all_locks": {},
	"release_lock": {},
	"repeat": {},
	"replace": {},
	"reverse": {},
	"right": {},
	"roles_graphml": {},
	"round": {},
	"row_count": {},
	"row_number": {},
	"rpad": {},
	"rtrim": {},
	"schema": {},
	"sec_to_time": {},
	"second": {},
	"session_user": {},
	"sha1": {},
	"sha": {},
	"sha2": {},
	"sign": {},
	"sin": {},
	"sleep": {},
	"soundex": {},
	"space": {},
	"sqrt": {},
	"st_area": {},
	"st_asbinary": {},
	"st_aswkb": {},
	"st_asgeojson": {},
	"st_astext": {},
	"st_aswkt": {},
	"st_buffer": {},
	"st_buffer_strategy": {},
	"st_centroid": {},
	"st_contains": {},
	"st_convexhull": {},
	"st_crosses": {},
	"st_difference": {},
	"st_dimension": {},
	"st_disjoint": {},
	"st_distance": {},
	"st_distance_sphere": {},
	"st_endpoint": {},
	"st_envelope": {},
	"st_equals": {},
	"st_exteriorring": {},
	"st_geohash": {},
	"st_geomcollfromtext": {},
	"st_geomcollfromtxt": {},
	"st_geomcollfromwkb": {},
	"st_geometrycollectionfromtext": {},
	"st_geometrycollectionfromwkb": {},
	"st_geometryn": {},
	"st_geometrytype": {},
	"st_geomfromgeojson": {},
	"st_geomfromtext": {},
	"st_geometryfromtext": {},
	"st_geomfromwkb": {},
	"st_geometryfromwkb": {},
	"st_interiorringn": {},
	"st_intersection": {},
	"st_intersects": {},
	"st_isclosed": {},
	"st_isempty": {},
	"st_issimple": {},
	"st_isvalid": {},
	"st_latfromgeohash": {},
	"st_latitude": {},
	"st_length": {},
	"st_linefromtext": {},
	"st_linestringfromtext": {},
	"st_linefromwkb": {},
	"st_linestringfromwkb": {},
	"st_longfromgeohash": {},
	"st_longitude": {},
	"st_makeenvelope": {},
	"st_mlinefromtext": {},
	"st_multilinestringfromtext": {},
	"st_mlinefromwkb": {},
	"st_multilinestringfromwkb": {},
	"st_mpointfromtext": {},
	"st_multipointfromtext": {},
	"st_mpointfromwkb": {},
	"st_multipointfromwkb": {},
	"st_mpolyfromtext": {},
	"st_multipolygonfromtext": {},
	"st_mpolyfromwkb": {},
	"st_multipolygonfromwkb": {},
	"st_numgeometries": {},
	"st_numinteriorring": {},
	"st_numinteriorrings": {},
	"st_numpoints": {},
	"st_overlaps": {},
	"st_pointfromgeohash": {},
	"st_pointfromtext": {},
	"st_pointfromwkb": {},
	"st_pointn": {},
	"st_polyfromtext": {},
	"st_polygonfromtext": {},
	"st_polyfromwkb": {},
	"st_polygonfromwkb": {},
	"st_simplify": {},
	"st_srid": {},
	"st_startpoint": {},
	"st_swapxy": {},
	"st_symdifference": {},
	"st_touches": {},
	"st_transform": {},
	"st_union": {},
	"st_validate": {},
	"st_within": {},
	"st_x": {},
	"st_y": {},
	"statement_digest": {},
	"statement_digest_text": {},
	"std": {},
	"stddev": {},
	"stddev_pop": {},
	"stddev_samp": {},
	"str_to_date": {},
	"strcmp": {},
	"subdate": {},
	"substr": {},
	"substring": {},
	"substring_index": {},
	"subtime": {},
	"sum": {},
	"sysdate": {},
	"system_user": {},
	"tan": {},
	"time": {},
	"time_format": {},
	"time_to_sec": {},
	"timediff": {},
	"timestamp": {},
	"timestampadd": {},
	"timestampdiff": {},
	"to_base64": {},
	"to_days": {},
	"to_seconds": {},
	"trim": {},
	"truncate": {},
	"ucase": {},
	"uncompress": {},
	"uncompressed_length": {},
	"unhex": {},
	"unix_timestamp": {},
	"updatexml": {},
	"upper": {},
	"user": {},
	"utc_date": {},
	"utc_time": {},
	"utc_timestamp": {},
	"uuid": {},
	"uuid_short": {},
	"uuid_to_bin": {},
	"validate_password_strength": {},
	"values": {},
	"var_pop": {},
	"var_samp": {},
	"variance": {},
	"version": {},
	"wait_for_executed_gtid_set": {},
	"wait_until_sql_thread_after_gtids": {},
	"week": {},
	"weekday": {},
	"weekofyear": {},
	"weight_string": {},
	"year": {},
	"yearweek": {},
}

func resolveColumnDefaults(ctx *sql.Context, a *Analyzer, n sql.Node, scope *Scope) (sql.Node, error) {
	span, _ := ctx.Span("resolveColumnDefaults")
	defer span.Finish()

	return plan.TransformUp(n, func(n sql.Node) (sql.Node, error) {
		switch node := n.(type) {
		case sql.SchemaModifiable:
			sch := node.Schema()
			var columns = make(map[string]indexedCol)
			for i, col := range sch {
				columns[strings.ToLower(col.Name)] = indexedCol{col, i}
			}
			newSch := make(sql.Schema, len(sch))

			for i, col := range sch {
				newCol := *col
				if col.Default == nil || col.Default.Resolved() {
					newSch[i] = &newCol
					continue
				}
				if sql.IsTextBlob(newCol.Type) && newCol.Default.IsLiteral() {
					return nil, sql.ErrInvalidTextBlobColumnDefault.New()
				}
				var err error
				sql.Inspect(newCol.Default.Expression, func(e sql.Expression) bool {
					switch expr := e.(type) {
					case sql.FunctionExpression:
						funcName := expr.FunctionName()
						if _, isValid := validColumnDefaultFuncs[funcName]; !isValid {
							err = sql.ErrInvalidColumnDefaultFunction.New(funcName, col.Name)
							return false
						}
						if (funcName == "now" || funcName == "current_timestamp") &&
							newCol.Default.IsLiteral() &&
							(!sql.IsTime(newCol.Type) || sql.Date == newCol.Type) {
							err = sql.ErrColumnDefaultDatetimeOnlyFunc.New()
							return false
						}
						return true
					case *plan.Subquery:
						err = sql.ErrColumnDefaultSubquery.New(col.Name)
						return false
					default:
						return true
					}
				})
				if err != nil {
					return nil, err
				}
				newCol.Default, err = sql.NewColumnDefaultValue(newCol.Default.Expression, newCol.Type, newCol.Default.IsLiteral(), newCol.Nullable)
				if err != nil {
					return nil, err
				}
				newSch[i] = &newCol
			}
			return node.WithSchema(newSch)
		default:
			return node, nil
		}
	})
}
