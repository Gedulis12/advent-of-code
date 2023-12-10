INPUT = './input'


def get_seed_range_dict(file):
    with open(INPUT, 'r') as f:
        lines = f.readlines()
        seed_list = [
                line.split(':')[1].strip().split(' ')[:] for
                line in
                lines if
                'seeds:' in
                line][0]
        f.close()

    chunked_list = list()
    chunk_size = 2

    for i in range(0, len(seed_list), chunk_size):
        chunked_list.append(seed_list[i:i+chunk_size])

    seed_dict = {}
    for i in range(len(chunked_list)):
        start = chunked_list[i][0]
        _range = chunked_list[i][1]
        end = int(start) + int(_range) - 1

        seed_dict[f'map_{i+1}'] = {}
        seed_dict[f'map_{i+1}']['start'] = int(start)
        seed_dict[f'map_{i+1}']['end'] = int(end)

    return seed_dict


with open(INPUT, 'r') as f:
    lines = f.readlines()
    seeds = [
            line.split(':')[1].strip().split(' ')[:] for
            line in
            lines if
            'seeds:' in
            line][0]

    for i in range(len(lines)):
        if 'seed-to-soil map:' in lines[i]:
            seed_to_soil = []
            j = 1
            while i+j < len(lines) and lines[i+j] != '\n':
                seed_to_soil.append(lines[i+j].strip().split(' '))
                j += 1
        elif 'soil-to-fertilizer map:' in lines[i]:
            soil_to_fert = []
            j = 1
            while i+j < len(lines) and lines[i+j] != '\n':
                soil_to_fert.append(lines[i+j].strip().split(' '))
                j += 1
        elif 'fertilizer-to-water map:' in lines[i]:
            fert_to_water = []
            j = 1
            while i+j < len(lines) and lines[i+j] != '\n':
                fert_to_water.append(lines[i+j].strip().split(' '))
                j += 1
        elif 'water-to-light map:' in lines[i]:
            water_to_light = []
            j = 1
            while i+j < len(lines) and lines[i+j] != '\n':
                water_to_light.append(lines[i+j].strip().split(' '))
                j += 1
        elif 'light-to-temperature map:' in lines[i]:
            light_to_temp = []
            j = 1
            while i+j < len(lines) and lines[i+j] != '\n':
                light_to_temp.append(lines[i+j].strip().split(' '))
                j += 1
        elif 'temperature-to-humidity map:' in lines[i]:
            temp_to_hum = []
            j = 1
            while i+j < len(lines) and lines[i+j] != '\n':
                temp_to_hum.append(lines[i+j].strip().split(' '))
                j += 1
        elif 'humidity-to-location map:' in lines[i]:
            hum_to_loc = []
            j = 1
            while i+j < len(lines) and lines[i+j] != '\n':
                hum_to_loc.append(lines[i+j].strip().split(' '))
                j += 1
    f.close()


def list_to_map(in_list):
    out_map = {
        f"map_{i}": {
            'dst_range_start': int(dst),
            'src_range_start': int(src),
            'range_len': int(length)
        }
        for i, (dst, src, length) in enumerate(in_list, start=1)
    }
    return out_map


seed_to_soil_map = list_to_map(seed_to_soil)
soil_to_fert_map = list_to_map(soil_to_fert)
fert_to_water_map = list_to_map(fert_to_water)
water_to_light_map = list_to_map(water_to_light)
light_to_temp_map = list_to_map(light_to_temp)
temp_to_hum_map = list_to_map(temp_to_hum)
hum_to_loc_map = list_to_map(hum_to_loc)


def get_mapping(input, map):
    for s in map:
        input = int(input)
        mapping = input
        for s in map:
            start = map[s]['src_range_start']
            end = start + map[s]['range_len']
            if input >= start and input <= end:
                mapping = input - start + map[s]['dst_range_start']
                return mapping
    return mapping


def get_location_from_seed(seed):
    soil = get_mapping(seed, seed_to_soil_map)
    fertilizer = get_mapping(soil, soil_to_fert_map)
    water = get_mapping(fertilizer, fert_to_water_map)
    light = get_mapping(water, water_to_light_map)
    temp = get_mapping(light, light_to_temp_map)
    hum = get_mapping(temp, temp_to_hum_map)
    location = get_mapping(hum, hum_to_loc_map)

    return location


def get_min_location(seeds):
    min_location = get_location_from_seed(seeds[0])

    for seed in seeds:
        location = get_location_from_seed(seed)

        if location < min_location:
            min_location = location

    return min_location


range_dict = get_seed_range_dict(INPUT)


def optimize_range_dict(tange_dict):
    for i in range_dict:
        for j in range_dict:
            if (
                    range_dict[i]['start'] > range_dict[j]['start']
                    and
                    range_dict[i]['start'] < range_dict[j]['end']
                    ):
                if range_dict[i]['end'] < range_dict[j]['end']:
                    range_dict[i]['start'] = range_dict[i]['end'] = 0
                elif range_dict[i]['end'] > range_dict[j]['end']:
                    range_dict[i]['start'] = range_dict[j]['end']
            if (
                    range_dict[i]['end'] > range_dict[j]['start']
                    and
                    range_dict[i]['end'] < range_dict[j]['end']
                    ):
                if range_dict[i]['start'] > range_dict[j]['start']:
                    range_dict[i]['end'] = range_dict[j]['start']
    return range_dict


def get_min_location_from_dict(range_dict):
    mins = []

    for i in range_dict:
        range_list = []
        start = range_dict[i]['start']
        end = range_dict[i]['end']
        idx = 0
        count = 1

        if end - start > 250000:

            for j in range(start, end):
                range_list.append(j)
                idx += 1

                if idx >= 250000:
                    print(f"checking map {i}, checked {count * 250000}/{end-start}")
                    count += 1
                    idx = 0
                    min_check = get_min_location(range_list)
                    mins.append(min_check)
                    range_list = []

            min_check = get_min_location(range_list)
            mins.append(min_check)

        else:

            for j in range(start, end):
                range_list.append(j)

            min_check = get_min_location(range_list)
            mins.append(min_check)
    return min(mins)



optimized_dict = optimize_range_dict(range_dict)
print(optimized_dict)
ans_2 = get_min_location_from_dict(optimized_dict)
print(ans_2)
