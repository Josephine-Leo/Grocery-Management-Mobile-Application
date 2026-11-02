MERGE INTO items AS t
USING stg_item_table AS s
ON t.name = s.title

-- If match update
WHEN MATCHED THEN
    UPDATE SET
        t.sub_cat = s.sub_category,
        t.price = s.price,
        t.discount = s.discount,
        t.rating = s.rating,
        t.currency = s.currency,
        t.feature = s.feature,
        t.description = s.product_description,
        t.title = s.title

-- Else insert
WHEN NOT MATCHED THEN
    INSERT (sub_cat, price, discount, rating, title, currency, feature, description)
    VALUES (s.sub_category, s.price, s.discount, s.rating, s.title, s.currency, s.feature, s.product_description);