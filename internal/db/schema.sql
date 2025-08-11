CREATE TABLE IF NOT EXISTS customers (
  id SERIAL PRIMARY KEY,
  code VARCHAR(20) UNIQUE NOT NULL,
  full_name TEXT NOT NULL,
  phone VARCHAR(30),
  email TEXT,
  address TEXT,
  date_of_birth DATE,
  driver_license VARCHAR(64),
  created_at TIMESTAMPTZ DEFAULT now(),
  updated_at TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE IF NOT EXISTS invoices (
  id SERIAL PRIMARY KEY,
  order_code VARCHAR(40) UNIQUE NOT NULL,
  customer_id INTEGER NOT NULL REFERENCES customers(id) ON DELETE CASCADE,
  items JSONB NOT NULL DEFAULT '[]',
  weight_lbs NUMERIC(10,2),
  subtotal NUMERIC(12,2) NOT NULL DEFAULT 0,
  surcharges JSONB NOT NULL DEFAULT '[]',
  total NUMERIC(12,2) NOT NULL DEFAULT 0,
  currency VARCHAR(10) NOT NULL DEFAULT 'USD',
  -- Money transfer fields (NEW)
  transfer_amount NUMERIC(12,2),     -- Số tiền chuyển
  transfer_fee NUMERIC(12,2),        -- Phí chuyển
  total_recipient NUMERIC(12,2),     -- Tổng tiền cho người nhận
  invoice_date DATE,                 -- Ngày hoá đơn
  bank_info_vn TEXT,                 -- Thông tin bank VN
  status VARCHAR(20) NOT NULL DEFAULT 'NEW',
  notes TEXT,
  created_at TIMESTAMPTZ DEFAULT now(),
  updated_at TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_invoices_customer_id ON invoices(customer_id);
