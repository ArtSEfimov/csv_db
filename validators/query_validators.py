import csv
import inspect
import os.path
import string


def file_exist_checking(filename):
    if not filename.endswith(".csv"):
        filename = f"{filename}.csv"
    if not os.path.isfile(filename):
        return False

    return True


def parameters_count_checking(filename, count):
    if not filename.endswith(".csv"):
        filename = f"{filename}.csv"
    with open(filename, 'r', encoding="utf-8") as file:
        csv_reader = csv.reader(file, delimiter='|')
        header = next(csv_reader)
        if count < len(header) - 1:
            return -1
        if count > len(header) - 1:
            return 1
        return 0


def check_id(id_value):
    called_function = inspect.stack()[1][3]
    allowed_chars = {"select_validator": "'*' or integer",
                     "update_validator": "integer",
                     "delete_validator": "integer"}
    id_value = id_value.strip(string.punctuation)
    try:
        id_to_int = int(id_value)
    except ValueError:
        return False, f"bad request parameter: ID must be {allowed_chars[called_function]}, your value: {id_value}"
    else:
        if id_to_int <= 0:
            return False, f"bad request parameter: ID must be positive, your value: {id_to_int}"
    return True, ""


def create_table_validator(query):
    query = query.split()
    if len(query) < 3:
        return False, "bad request: too little query parameters"

    return True, ""


def select_validator(query):
    query = query.split()
    request_type, table_name, request_parameter, *_ = query

    if not file_exist_checking(table_name):
        return False, f"the file '{table_name}' does not exist"

    if len(query) < 3:
        return False, "bad request: too little query parameters"
    if len(query) > 3:
        return False, "bad request: too many query parameters"

    if request_parameter == '*':
        return True, ""

    return check_id(request_parameter)


def insert_validator(query):
    query = query.split()
    if len(query) < 3:
        return False, "bad request: too little query parameters"

    request_type, table_name, *request_parameters = query

    if not file_exist_checking(table_name):
        return False, f"the file '{table_name}' does not exist"

    parameters_count = parameters_count_checking(table_name, len(request_parameters))

    if parameters_count > 0:
        return False, "too many fields"
    if parameters_count < 0:
        return False, "too little fields"

    return True, ""


def update_validator(query):
    query = query.split()
    if len(query) < 4:
        return False, "bad request: too little query parameters"

    request_type, table_name, record_id, *request_parameters = query

    if not file_exist_checking(table_name):
        return False, f"the file '{table_name}' does not exist"

    parameters_count = parameters_count_checking(table_name, len(request_parameters))
    if parameters_count > 0:
        return False, "too many fields"
    if parameters_count < 0:
        return False, "too little fields"

    return check_id(record_id)


def delete_validator(query):
    query = query.split()

    request_type, table_name, record_id, *_ = query

    if not file_exist_checking(table_name):
        return False, f"the file '{table_name}' does not exist"

    if len(query) < 3:
        return False, "bad request: too little query parameters"
    if len(query) > 3:
        return False, "bad request: too many query parameters"

    return check_id(record_id)
