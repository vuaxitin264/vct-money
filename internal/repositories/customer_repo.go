package repositories

import (
    "context"
    "database/sql"
    "vch/internal/models"
)

type CustomerRepo struct{ DB *sql.DB }

func (r *CustomerRepo) Create(ctx context.Context, c *models.Customer) error {
    return r.DB.QueryRowContext(ctx,
        `INSERT INTO customers(code,full_name,phone,email,address,date_of_birth,driver_license) VALUES($1,$2,$3,$4,$5,$6,$7) RETURNING id`,
        c.Code, c.FullName, c.Phone, c.Email, c.Address, c.DateOfBirth, c.DriverLicense,
    ).Scan(&c.ID)
}

func (r *CustomerRepo) GetByID(ctx context.Context, id int64) (*models.Customer, error) {
    row := r.DB.QueryRowContext(ctx, `SELECT id,code,full_name,phone,email,address,date_of_birth,driver_license,created_at,updated_at FROM customers WHERE id=$1`, id)
    var c models.Customer
    if err := row.Scan(&c.ID,&c.Code,&c.FullName,&c.Phone,&c.Email,&c.Address,&c.DateOfBirth,&c.DriverLicense,&c.CreatedAt,&c.UpdatedAt); err != nil { return nil, err }
    return &c, nil
}

func (r *CustomerRepo) List(ctx context.Context, q string, limit, offset int) ([]models.Customer, error) {
    if limit<=0 { limit=50 }
    rows, err := r.DB.QueryContext(ctx,
      `SELECT id,code,full_name,phone,email,address,date_of_birth,driver_license,created_at,updated_at
       FROM customers
       WHERE ($1='' OR full_name ILIKE '%'||$1||'%' OR phone ILIKE '%'||$1||'%')
       ORDER BY id DESC LIMIT $2 OFFSET $3`, q, limit, offset)
    if err != nil { return nil, err }
    defer rows.Close()
    var out []models.Customer
    for rows.Next(){
        var c models.Customer
        if err := rows.Scan(&c.ID,&c.Code,&c.FullName,&c.Phone,&c.Email,&c.Address,&c.DateOfBirth,&c.DriverLicense,&c.CreatedAt,&c.UpdatedAt); err!=nil { return nil, err }
        out = append(out, c)
    }
    return out, nil
}
