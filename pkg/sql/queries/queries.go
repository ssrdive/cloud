package queries

const ALL_ITEMS = `
	SELECT id, item_id, model_id, item_category_id, page_no, item_no, foreign_id, item_name, price FROM item
`

const ITEM_DETAILS_BY_ITEM_ID = `
	SELECT I.id, I.item_id, I.model_id, M.name AS model_name, I.item_category_id, IC.name AS item_category_name, I.page_no, I.item_no, I.foreign_id, I.item_name, I.price
	FROM item I 
	LEFT JOIN model M ON M.id = I.model_id
	LEFT JOIN item_category IC ON IC.id = I.item_category_id
	WHERE I.item_id = ?
`

const ITEM_DETAILS_BY_ID = `
	SELECT I.id, I.item_id, I.model_id, M.name AS model_name, I.item_category_id, IC.name AS item_category_name, I.page_no, I.item_no, I.foreign_id, I.item_name, I.price
	FROM item I 
	LEFT JOIN model M ON M.id = I.model_id
	LEFT JOIN item_category IC ON IC.id = I.item_category_id
	WHERE I.id = ?
`

const SEARCH_ITEMS = `
	SELECT id, item_id, model_id, item_category_id, page_no, item_no, foreign_id, item_name, price 
	FROM item
	WHERE (? IS NULL OR CONCAT(item_id, foreign_id, item_name) LIKE ?)
`

const SALE_WATCHES = `
	SELECT  SW.id, SW.sale_id, SW.content, SW.date, SW.due_date, SW.user, SW.closed, SW.closed_by, DATEDIFF(NOW(), due_date) as expires
	FROM sale_watch SW
	LEFT JOIN user U 
	ON SW.closed_by = U.username
	ORDER BY expires
`

const SALE_ALL = `
	SELECT S.id, COUNT(SC.sale_id) as sale_comments,  R.name as region_name, S.date, S.sys_date,  S.chassis_no, S.customer_name, S.customer_address, S.customer_contact,  S.invoice_no, S.price,  S.institute, S.advance, S.deleted
	FROM sale S
	LEFT JOIN sale_comment SC ON SC.sale_id = S.id
	LEFT JOIN region R ON S.region = R.id
	WHERE S.deleted = 0
	GROUP BY S.id, date, sys_date
	ORDER BY S.id DESC, date
`

const SALE_ALL_DUMMY = `
	SELECT S.id, COUNT(SC.sale_id) as sale_comments, R.name as region_name, T.name as territory_name,  S.location,  D.name as sd_location, S.location_fk,  S.chassis_no, S.customer_name, S.customer_address, S.customer_contact, M.name as model_name, S.invoice_no, S.price, ST.name as sale_type_name, S.institute, S.advance, S.date, S.sys_date
	FROM sale S
	LEFT JOIN dealer D ON S.location_fk = D.id
	AND S.location_fk <> 0
	LEFT JOIN region R ON S.region = R.id
	LEFT JOIN territory T ON S.territory = T.id
	LEFT JOIN model M ON S.model =M.id
	LEFT JOIN sale_type ST ON S.sale_type = ST.id
	LEFT JOIN sale_comment SC ON SC.sale_id = S.id
	LEFT JOIN sale_completed_type SCT ON SCT.id = S.sale_completed_type_id
	WHERE S.deleted = 0
	GROUP BY S.id, region_name, territory_name, date, sys_date
	ORDER BY S.id DESC, date
`
const SEARCH_CLOUDIDINFO = `
	SELECT S.id, COUNT(SC.sale_id) as sale_comments, S.chassis_no, U.name as officer_name,  R.name as region_name, T.name as territory_name, S.date, S.sys_date, S.location,  D.name as sd_location, S.customer_name,  M.name as model_name, S.location_fk, S.deleted
	FROM sale S
	LEFT JOIN user U ON S.officer = U.username 
	LEFT JOIN region R ON S.region = R.id 
	LEFT JOIN territory T ON S.territory = T.id 
	LEFT JOIN model M on S.model = M.id 
	LEFT JOIN sale_comment SC ON SC.sale_id = S.id
	LEFT JOIN dealer D ON S.location_fk = D.id AND S.location_fk <> 0 
	WHERE  S.deleted = 0 AND (? IS NULL OR CONCAT(S.id) LIKE ?)
	GROUP BY S.id
`
const SEARCH_CHASSISINFO = `
	SELECT S.id, COUNT(SC.sale_id) as sale_comments, S.chassis_no, U.name as officer_name,  R.name as region_name, T.name as territory_name, S.date, S.sys_date, S.location,  D.name as sd_location, S.customer_name,  M.name as model_name, S.location_fk, S.deleted
	FROM sale S
	LEFT JOIN user U ON S.officer = U.username 
	LEFT JOIN region R ON S.region = R.id 
	LEFT JOIN territory T ON S.territory = T.id 
	LEFT JOIN model M on S.model = M.id 
	LEFT JOIN sale_comment SC ON SC.sale_id = S.id
	LEFT JOIN dealer D ON S.location_fk = D.id AND S.location_fk <> 0 
	WHERE  S.deleted = 0 AND (? IS NULL OR CONCAT(S.chassis_no) LIKE ?)
	GROUP BY S.id, S.chassis_no, officer_name,  region_name, territory_name, S.date, S.sys_date, sd_location, S.customer_name, model_name
`

const SEARCH_SALEINFO = `
	SELECT S.id, COUNT(SC.sale_id) as sale_comments, S.chassis_no, U.name as officer_name,  R.name as region_name, T.name as territory_name, S.date, S.sys_date, S.location,  D.name as sd_location, S.customer_name,  M.name as model_name, S.location_fk, S.deleted
	FROM sale S
	LEFT JOIN user U ON S.officer = U.username 
	LEFT JOIN region R ON S.region = R.id 
	LEFT JOIN territory T ON S.territory = T.id 
	LEFT JOIN model M on S.model = M.id 
	LEFT JOIN sale_comment SC ON SC.sale_id = S.id
	LEFT JOIN dealer D ON S.location_fk = D.id AND S.location_fk <> 0 
	WHERE  S.deleted = 0 AND (? IS NULL OR CONCAT(S.chassis_no, U.name, R.name, T.name, S.date, S.sys_date, S.location, D.name, S.chassis_no, S.customer_name, M.name) LIKE ?)
	GROUP BY S.id, officer_name, region_name, territory_name, date, sys_date, S.location, sd_location, S.location_fk, S.chassis_no, S.customer_name, S.customer_address, S.customer_contact, model_name
	ORDER BY S.id DESC
`

const CLOUD_ID_INFO = `
	SELECT S.id,  U.name as officer_name, R.name as region_name, T.name as territory_name, S.location, D.name as sd_location, DT.name as dealer_territory, S.location_fk, S.advance, S.date, S.chassis_no, S.customer_name, S.customer_address, S.customer_contact, M.name as model_name, ST.name as sale_type_name, S.invoice_no, S.price,  S.institute, S.deleted, S.region, S.territory, D.territory_id, S.model, S.sale_type, S.officer, S.latitude, S.longitude, S.verified, S.verified_by as verified_by , S.verified_on as verified_on, S.sale_completed, SCT.name as sale_completed_type_name, 
	sale_completed_remarks, sale_completed_by, sale_completed_on , S.commision_paid, S.commission_paid_marked_by, S.commission_paid_marked_on, S.commission_paid_remarks, SC.sale_id, S.sale_completed_type_id 
	FROM sale S
	LEFT JOIN dealer D ON S.location_fk = D.id AND S.location_fk <> 0 
	LEFT JOIN user U ON S.officer = U.username  
	LEFT JOIN region R ON S.region = R.id 
	LEFT JOIN territory T ON S.territory = T.id 
	LEFT JOIN model M on S.model = M.id 
	LEFT JOIN territory DT ON D.territory_id = DT.id 
	LEFT JOIN sale_type ST ON S.sale_type = ST.id 
	LEFT JOIN sale_comment SC ON SC.sale_id = S.id 
	LEFT JOIN sale_completed_type SCT ON SCT.id = S.sale_completed_type_id 
	WHERE S.deleted = 0 and S.id= ? 
`

const SALE_VERIFY = `
	UPDATE sale 
	SET verified = 1, verified_by = ?, verified_on = ? WHERE id = ?
`

const COMMENTS = `
	SELECT SC.username, U.name, SC.date, SC.text, SC.attachment, sale_id
	FROM sale_comment SC 
	LEFT JOIN user U ON SC.username = U.username 
	WHERE sale_id = ? 
	ORDER BY SC.date DESC 
`

const MARK_SALE_COMPLETE = `
	SELECT id, name 
	FROM sale_completed_type
`
