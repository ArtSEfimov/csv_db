import sys

import query_validators

validators = {
    "CREATE_TABLE": query_validators.create_table_validator,
    "SELECT": query_validators.select_validator,
    "INSERT": query_validators.insert_validator,
    "UPDATE": query_validators.update_validator,
    "DELETE": query_validators.delete_validator,
}

query = sys.argv[1]
request_type = query.split()[0]
result, error = validators[request_type](query)
print(result, error, sep='\n', end='')
