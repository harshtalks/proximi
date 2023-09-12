# Businesses API

Use this API to access businesses nearby you, based on latitude and longitude you provide.

## Params

### 1. Lat

- Latitude of the location of the user (string value)
- Required

### 2. Long

- Longitude of the location of the user (string value)
- Required

### 3. range

- Range of the result
- You can specify the range in kilometers
- default is the area of radius .5kms to 2kms
- optional

### 4. page

- This API is paginated.
- default page number is 1
- optional

### 5. perPage

- per page results
- default is 10
- max is 100
- minimum is 10
- optional
