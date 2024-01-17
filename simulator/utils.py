from calculator import Tick

def fe(num):
    return '{:.18f}'.format(num)

def distinct(ticks: [Tick]):
    unique_values_set = set()
    filtered_list = [obj for obj in ticks if obj.rate not in unique_values_set and not unique_values_set.add(obj.rate)]
    return filtered_list