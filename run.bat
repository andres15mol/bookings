go build -o bookings.exe ./cmd/web/. || exit /b
bookings.exe -dbname=bookings -dbuser=postgres -dbpass=dog1234 -cache=false -production=false 