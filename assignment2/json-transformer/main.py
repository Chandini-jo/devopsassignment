import json
import re
from datetime import datetime

def sanitize_value(value):
    """Sanitize the value of trailing and leading whitespace."""
    return value.strip()

def transform_string(value):
    """Transform value to the String data type."""
    sanitized_value = sanitize_value(value)
    # Transform RFC3339 formatted Strings to Unix Epoch in Numeric data type.
    if re.match(r'\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}Z', sanitized_value):
        return int(datetime.strptime(sanitized_value, '%Y-%m-%dT%H:%M:%SZ').timestamp())
    # Omit fields with empty values.
    elif sanitized_value == '':
        return None
    else:
        return sanitized_value

def transform_number(value):
    """Transform value to the Numeric data type."""
    sanitized_value = sanitize_value(value)
    # Omit fields with invalid Numeric values.
    try:
        # Strip the leading zeros.
        return float(sanitized_value) if '.' in sanitized_value else int(sanitized_value)
    except ValueError:
        return None

def transform_boolean(value):
    """Transform value to the Boolean data type."""
    sanitized_value = sanitize_value(value)
    # Omit fields with invalid Boolean values.
    if sanitized_value in {'1', 't', 'T', 'TRUE', 'true', 'True'}:
        return True
    elif sanitized_value in {'0', 'f', 'F', 'FALSE', 'false', 'False'}:
        return False
    else:
        return None

def transform_null(value):
    """Represent a null literal when the value is true."""
    sanitized_value = sanitize_value(value)
    # Omit fields with invalid Boolean values.
    if sanitized_value in {'1', 't', 'T', 'TRUE', 'true', 'True'}:
        return None
    elif sanitized_value in {'0', 'f', 'F', 'FALSE', 'false', 'False'}:
        return False
    else:
        return None

def transform_list(value):
    """Transform value to the List data type."""
    if not isinstance(value, list):
        return None
    result = []
    for item in value:
        if not isinstance(item, dict):
            continue
        for data_type, val in item.items():
            if data_type == 'S':
                transformed_value = transform_string(val)
            elif data_type == 'N':
                transformed_value = transform_number(val)
            elif data_type == 'BOOL':
                transformed_value = transform_boolean(val)
            elif data_type == 'NULL':
                transformed_value = transform_null(val)
            else:
                continue
            if transformed_value is not None:
                result.append(transformed_value)
    return result if result else None

def transform_map(value):
    """Transform value to the Map data type."""
    result = {}
    for key, item in sorted(value.items()):
        for data_type, val in item.items():
            if data_type == 'S':
                transformed_value = transform_string(val)
            elif data_type == 'N':
                transformed_value = transform_number(val)
            elif data_type == 'BOOL':
                transformed_value = transform_boolean(val)
            elif data_type == 'NULL':
                transformed_value = transform_null(val)
            elif data_type == 'L':
                transformed_value = transform_list(val)
            elif data_type == 'M':
                transformed_value = transform_map(val)
            else:
                continue
            if transformed_value is not None:
                result[key] = transformed_value
    return result if result else None

def json_transformer(input_json):
    result = {}
    for key, item in input_json.items():
        sanitized_key = sanitize_value(key)
        if sanitized_key:  # Omit fields with empty keys.
            for data_type, val in item.items():
                if data_type == 'S':
                    transformed_value = transform_string(val)
                elif data_type == 'N':
                    transformed_value = transform_number(val)
                elif data_type == 'BOOL':
                    transformed_value = transform_boolean(val)
                elif data_type == 'NULL':
                    transformed_value = transform_null(val)
                elif data_type == 'L':
                    transformed_value = transform_list(val)
                elif data_type == 'M':
                    transformed_value = transform_map(val)
                else:
                    continue
                if transformed_value is not None:
                    result[sanitized_key] = transformed_value
    return [result] if result else None


#Specify the file path
file_path = "input.json"

#Loading the contents of json file
with open(file_path, "r") as file:
    json_dict = json.load(file)

# Transforming the JSON input
transformed_json = json_transformer(json_dict)

# Print the transformed JSON
print(json.dumps(transformed_json, indent=2))