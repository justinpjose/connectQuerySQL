package backbonev2

import "fmt"

var (
	getProductCodesInjector = `
	SELECT supref ->> 'productCode' pcode
	FROM
	(
		SELECT
			jsonb_array_elements(sup -> 'supplierReferences') supref
		FROM
			(
				SELECT
					jsonb_array_elements(opt -> 'suppliers') sup
				FROM
					(
						SELECT
							jsonb_array_elements(style -> 'options') opt
						FROM
							injector
					) as opts
			) as sups
	) as suprefs
	`
)

func getNonEmptyProductCodesInjector() string {
	query := fmt.Sprintf(`
	SELECT pcode
	from (
		%s
		) as pcodes
	where pcode != ''
	`, getProductCodesInjector)

	return query
}

func GetProductCodesInODBMSNotInjectorQuery() string {
	query := fmt.Sprintf(`
	SELECT id
	FROM odbms
	EXCEPT
	%s
	`, getNonEmptyProductCodesInjector())

	return query
}
