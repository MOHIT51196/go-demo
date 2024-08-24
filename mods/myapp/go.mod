module myapp

go 1.22.6

replace example.com/db => ../db

require example.com/db v0.0.0-00010101000000-000000000000 // indirect
