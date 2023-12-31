INPUT = './input'


def get_seed_range_list(file):
    with open(INPUT, 'r') as f:
        lines = f.readlines()
        seed_list = [
                line.split(':')[1].strip().split(' ')[:] for
                line in
                lines if
                'seeds:' in
                line][0]
        f.close()

    range_list = list()
    chunk_size = 2

    for i in range(0, len(seed_list), chunk_size):
        start = int(seed_list[i])
        end = int(seed_list[i]) + int(seed_list[i+1]) - 1
        seed_range = (start, end)
        range_list.append(seed_range)

    return range_list


with open(INPUT, 'r') as f:
    lines = f.readlines()
    seeds = [
            line.split(':')[1].strip().split(' ')[:] for
            line in
            lines if
            'seeds:' in
            line][0]

    mappings = []
    for i in range(len(lines)):
        if 'seed-to-soil map:' in lines[i]:
            seed_to_soil = []
            j = 1
            while i+j < len(lines) and lines[i+j] != '\n':
                dst_start = lines[i+j].strip().split(' ')[0]
                src_start = lines[i+j].strip().split(' ')[1]
                range_len = lines[i+j].strip().split(' ')[2]
                mapping_tuple = (dst_start, src_start, range_len)
                seed_to_soil.append(mapping_tuple)
                j += 1
            mappings.append(seed_to_soil)
        elif 'soil-to-fertilizer map:' in lines[i]:
            soil_to_fert = []
            j = 1
            while i+j < len(lines) and lines[i+j] != '\n':
                dst_start = lines[i+j].strip().split(' ')[0]
                src_start = lines[i+j].strip().split(' ')[1]
                range_len = lines[i+j].strip().split(' ')[2]
                mapping_tuple = (dst_start, src_start, range_len)
                soil_to_fert.append(mapping_tuple)
                j += 1
            mappings.append(soil_to_fert)
        elif 'fertilizer-to-water map:' in lines[i]:
            fert_to_water = []
            j = 1
            while i+j < len(lines) and lines[i+j] != '\n':
                dst_start = lines[i+j].strip().split(' ')[0]
                src_start = lines[i+j].strip().split(' ')[1]
                range_len = lines[i+j].strip().split(' ')[2]
                mapping_tuple = (dst_start, src_start, range_len)
                fert_to_water.append(mapping_tuple)
                j += 1
            mappings.append(fert_to_water)
        elif 'water-to-light map:' in lines[i]:
            water_to_light = []
            j = 1
            while i+j < len(lines) and lines[i+j] != '\n':
                dst_start = lines[i+j].strip().split(' ')[0]
                src_start = lines[i+j].strip().split(' ')[1]
                range_len = lines[i+j].strip().split(' ')[2]
                mapping_tuple = (dst_start, src_start, range_len)
                water_to_light.append(mapping_tuple)
                j += 1
            mappings.append(water_to_light)
        elif 'light-to-temperature map:' in lines[i]:
            light_to_temp = []
            j = 1
            while i+j < len(lines) and lines[i+j] != '\n':
                dst_start = lines[i+j].strip().split(' ')[0]
                src_start = lines[i+j].strip().split(' ')[1]
                range_len = lines[i+j].strip().split(' ')[2]
                mapping_tuple = (dst_start, src_start, range_len)
                light_to_temp.append(mapping_tuple)
                j += 1
            mappings.append(light_to_temp)
        elif 'temperature-to-humidity map:' in lines[i]:
            temp_to_hum = []
            j = 1
            while i+j < len(lines) and lines[i+j] != '\n':
                dst_start = lines[i+j].strip().split(' ')[0]
                src_start = lines[i+j].strip().split(' ')[1]
                range_len = lines[i+j].strip().split(' ')[2]
                mapping_tuple = (dst_start, src_start, range_len)
                temp_to_hum.append(mapping_tuple)
                j += 1
            mappings.append(temp_to_hum)
        elif 'humidity-to-location map:' in lines[i]:
            hum_to_loc = []
            j = 1
            while i+j < len(lines) and lines[i+j] != '\n':
                dst_start = lines[i+j].strip().split(' ')[0]
                src_start = lines[i+j].strip().split(' ')[1]
                range_len = lines[i+j].strip().split(' ')[2]
                mapping_tuple = (dst_start, src_start, range_len)
                hum_to_loc.append(mapping_tuple)
                j += 1
            mappings.append(hum_to_loc)
    f.close()


def get_location_from_seed(seed, mappings):
    for mapping in mappings:
        for map in mapping:
            dst_start, src_start, range_len = map
            src_end = int(src_start) + int(range_len) - 1
            if (int(src_start) <= int(seed) <= int(src_end)):
                seed = int(seed) - int(src_start) + int(dst_start)
                break
    return seed


def get_min_location(seeds, mappings):
    min_location = get_location_from_seed(seeds[0], mappings)

    for seed in seeds:
        location = get_location_from_seed(seed, mappings)

        if location < min_location:
            min_location = location

    return min_location


seeds = get_seed_range_list(INPUT)


def is_all_range_unmapped(r, map):
    range_start, range_end = r
    range_start = int(range_start)
    range_end = int(range_end)
    dst_start, src_start, range_len = map
    dst_start = int(dst_start)
    src_start = int(src_start)
    src_end = int(src_start) + int(range_len) - 1
    src_end = int(src_end)
    dst_end = int(dst_start) + int(range_len) - 1
    dst_end = int(dst_end)

    if ((range_start < src_start and range_end < src_start) or
            (range_end > src_end and range_start > src_end)):
        curr_range = (range_start, range_end)
        return curr_range


def is_all_range_mapped(r, map):
    range_start, range_end = r
    range_start = int(range_start)
    range_end = int(range_end)
    dst_start, src_start, range_len = map
    dst_start = int(dst_start)
    src_start = int(src_start)
    src_end = int(src_start) + int(range_len) - 1
    src_end = int(src_end)
    dst_end = int(dst_start) + int(range_len) - 1
    dst_end = int(dst_end)

    if (src_start <= range_start <= src_end and
            src_start <= range_end <= src_end):
        mapped_range_start = range_start - src_start + dst_start
        mapped_range_end = range_end - src_start + dst_start
        curr_range = (mapped_range_start, mapped_range_end)
        return curr_range


def is_start_range_mapped(r, map, only_mapped=True):
    range_start, range_end = r
    range_start = int(range_start)
    range_end = int(range_end)
    dst_start, src_start, range_len = map
    dst_start = int(dst_start)
    src_start = int(src_start)
    src_end = int(src_start) + int(range_len) - 1
    src_end = int(src_end)
    dst_end = int(dst_start) + int(range_len) - 1
    dst_end = int(dst_end)

    if (range_start >= src_start and
            range_start <= src_end and
            range_end > src_end):

        mapped_range_start = range_start - src_start + dst_start
        mapped_range_end = dst_end
        range_1 = (mapped_range_start, mapped_range_end)

        unmapped_range_start = src_end + 1
        unmapped_range_end = range_end
        range_2 = (unmapped_range_start, unmapped_range_end)

        if only_mapped:
            return range_1
        else:
            return range_2


def is_end_range_mapped(r, map, only_mapped=True):
    range_start, range_end = r
    range_start = int(range_start)
    range_end = int(range_end)
    dst_start, src_start, range_len = map
    dst_start = int(dst_start)
    src_start = int(src_start)
    src_end = int(src_start) + int(range_len) - 1
    src_end = int(src_end)
    dst_end = int(dst_start) + int(range_len) - 1
    dst_end = int(dst_end)

    if (range_start < src_start and
            range_end >= src_start and
            range_end < src_end):
        unmapped_range_start = range_start
        unmapped_range_end = src_start - 1
        range_1 = (unmapped_range_start, unmapped_range_end)

        mapped_range_start = dst_start
        mapped_range_end = range_end - src_start + dst_start
        range_2 = (mapped_range_start, mapped_range_end)

        if only_mapped:
            return range_2
        else:
            return range_1


def split_range_by_mapping(seed_range, mapping):
    ranges = list()

    unmapped = list()
    for map in mapping:
       #print(f'current map: {map}')

        # all range within mapping
        curr_range = is_all_range_mapped(seed_range, map)
        if curr_range:
            if (curr_range not in ranges):
                ranges.append(curr_range)
               #print(f'all range within mapping, produced: {curr_range}')

        # only end of the range included in mapping
        curr_range = is_end_range_mapped(seed_range, map)
        curr_range_unmapped = is_end_range_mapped(seed_range, map, False)
        if curr_range:
            if (curr_range not in ranges):
                ranges.append(curr_range)
               #print(f'only end of the range in mapping, produced: {curr_range}')

        if curr_range_unmapped:
            if (curr_range_unmapped not in unmapped):
                unmapped.append(curr_range_unmapped)

        # only start of the range included in mapping
        curr_range = is_start_range_mapped(seed_range, map)
        curr_range_unmapped = is_start_range_mapped(seed_range, map, False)
        if curr_range:
            if (curr_range not in ranges):
                ranges.append(curr_range)
               #print(f'only start of the range in mapping, produced: {curr_range}')

        if curr_range_unmapped:
            if (curr_range_unmapped not in unmapped):
                unmapped.append(curr_range_unmapped)

        # range outside of mapping
        curr_range = is_all_range_unmapped(seed_range, map)
        if curr_range:
            if (curr_range not in ranges):
                if(curr_range == seed_range and len(ranges) > 0):
                    pass
                else:
                    ranges.append(curr_range)
                   #print(f'all range outside mapping, produced: {curr_range}')

        #print(f'unmapped before: {unmapped}')
        for um in unmapped:
            #print(f'{um}')
            for map in mapping:
                #print(f'\t checking unmapped range: {um}, current map: {map}')
                curr_range = is_all_range_mapped(um, map)
                if curr_range:
                    #print(f'\t deleting: {um} from: {unmapped}')
                    unmapped.remove(um)
        #print(f'unmapped after: {unmapped}')

        for uml in unmapped:
            if uml not in ranges:
                ranges.append(uml)

    if len(ranges) > 1 and ranges[0] == (seed_range):
        return ranges[1:]
    else:
        return ranges


def get_seeds_to_check(seeds_ranges, mappings):
    min_seeds = list()
    ranges = list()
    for r in seeds_ranges:
        new_ranges = [r]
        for mapping in mappings:
           #print(f'START: mapping ranges: {new_ranges}')
            temp = list()
            for nr in new_ranges:
               #print(f'mapping range: {nr}')
                new_range = split_range_by_mapping(nr, mapping)
                for n in new_range:
                    temp.append(n)
               #print(f'produced new range: {new_range}')
            new_ranges = []
            for t in temp:
                new_ranges.append(t)
           #print(f'DONE: produced new ranges: {new_ranges}\n')
        ranges.append(new_ranges)

        for r in ranges:
            for range_tuple in r:
                min_seeds.append(range_tuple[0])

   #print(ranges)
   #print(min_seeds)
    return min_seeds


min_seeds = get_seeds_to_check(seeds, mappings)
min = min(min_seeds)
print(min)
