import csv
import os

violations = [
    {'id': 1, 'name':	'TOTAL Movers'},
    {'id': 2, 'name':	'Backing Unsafely'},
    {'id': 3, 'name':	'Bike Lane'},
    {'id': 5, 'name':	'Bus Lane'},
    {'id': 7, 'name':	'Cell Phone'},
    {'id': 8, 'name':	'Commercial Veh on Pkwy'},
    {'id': 9, 'name':	'Cruising For Passengers'},
    {'id': 13, 'name':	'Disobey Traffic Control Device'},
    {'id': 14, 'name':	'Driving Too Slow'},
    {'id': 15, 'name':	'Equipment'},
    {'id': 16, 'name':	'Equipment (Other)'},
    {'id': 17, 'name':	'Excessive Noise'},
    {'id': 18, 'name':	'Fail to Keep Right'},
    {'id': 21, 'name':	'Fail to yield Right of Way to Pedestrian'},
    {'id': 22, 'name':	'Failure to Signal'},
    {'id': 23, 'name':	'Failure to yield Right of Way to Vehicle'},
    {'id': 24, 'name':	'Following Too Closely'},
    {'id': 27, 'name':	'Improper Passing'},
    {'id': 28, 'name':	'Improper Taxi Pickup'},
    {'id': 29, 'name':	'Improper Turn'},
    {'id': 30, 'name':	'Improper/Missing Plates'},
    {'id': 31, 'name':	'Lamps and Other Equipment on Bicycle'},
    {'id': 32, 'name':	'Motorcycle (Other)'},
    {'id': 34, 'name':	'Not Giving R of W to Veh.'},
    {'id': 36, 'name':	'Obstructed Plate'},
    {'id': 37, 'name':	'One Way Street'},
    {'id': 38, 'name':	'Other Movers'},
    {'id': 39, 'name':	'Overheight'},
    {'id': 40, 'name':	'Overlength'},
    {'id': 41, 'name':	'Overweight'},
    {'id': 42, 'name':	'Overwidth'},
    {'id': 43, 'name':	'Pavement Markings'},
    {'id': 44, 'name':	'Red Light'},
    {'id': 46, 'name':	'School Bus'},
    {'id': 48, 'name':	'Scooter In NYC'},
    {'id': 49, 'name':	'Seat Belt'},
    {'id': 50, 'name':	'Speeding'},
    {'id': 51, 'name':	'Spillback'},
    {'id': 52, 'name':	'TBTA Rule'},
    {'id': 53, 'name':	'TLC (Other)'},
    {'id': 55, 'name':	'Tints'},
    {'id': 56, 'name':	'Truck Route'},
    {'id': 59, 'name':	'Uninspected'},
    {'id': 60, 'name':	'Uninsured'},
    {'id': 61, 'name':	'Unlicensed Operator'},
    {'id': 62, 'name':	'Unregistered'},
    {'id': 63, 'name':	'Unsafe Lane Change'}
]

precincts = [
    {'id': 1, 'name':	'city'},
    {'id': 2, 'name':	'001'},
    {'id': 3, 'name':	'005'},
    {'id': 4, 'name':	'006'},
    {'id': 5, 'name':	'007'},
    {'id': 6, 'name':	'009'},
    {'id': 7, 'name':	'010'},
    {'id': 8, 'name':	'013'},
    {'id': 9, 'name':	'014'},
    {'id': 10, 'name':	'017'},
    {'id': 11, 'name':	'018'},
    {'id': 12, 'name':	'019'},
    {'id': 13, 'name':	'020'},
    {'id': 14, 'name':	'022'},
    {'id': 15, 'name':	'023'},
    {'id': 16, 'name':	'024'},
    {'id': 17, 'name':	'025'},
    {'id': 18, 'name':	'026'},
    {'id': 19, 'name':	'028'},
    {'id': 20, 'name':	'030'},
    {'id': 21, 'name':	'032'},
    {'id': 22, 'name':	'033'},
    {'id': 23, 'name':	'034'},
    {'id': 24, 'name':	'040'},
    {'id': 25, 'name':	'041'},
    {'id': 26, 'name':	'042'},
    {'id': 27, 'name':	'043'},
    {'id': 28, 'name':	'044'},
    {'id': 29, 'name':	'045'},
    {'id': 30, 'name':	'046'},
    {'id': 31, 'name':	'047'},
    {'id': 32, 'name':	'048'},
    {'id': 33, 'name':	'049'},
    {'id': 34, 'name':	'050'},
    {'id': 35, 'name':	'052'},
    {'id': 36, 'name':	'060'},
    {'id': 37, 'name':	'061'},
    {'id': 38, 'name':	'062'},
    {'id': 39, 'name':	'063'},
    {'id': 40, 'name':	'066'},
    {'id': 41, 'name':	'067'},
    {'id': 42, 'name':	'068'},
    {'id': 43, 'name':	'069'},
    {'id': 44, 'name':	'070'},
    {'id': 45, 'name':	'071'},
    {'id': 46, 'name':	'072'},
    {'id': 47, 'name':	'073'},
    {'id': 48, 'name':	'075'},
    {'id': 49, 'name':	'076'},
    {'id': 50, 'name':	'077'},
    {'id': 51, 'name':	'078'},
    {'id': 52, 'name':	'079'},
    {'id': 53, 'name':	'081'},
    {'id': 54, 'name':	'083'},
    {'id': 55, 'name':	'084'},
    {'id': 56, 'name':	'088'},
    {'id': 57, 'name':	'090'},
    {'id': 58, 'name':	'094'},
    {'id': 59, 'name':	'100'},
    {'id': 60, 'name':	'101'},
    {'id': 61, 'name':	'102'},
    {'id': 62, 'name':	'103'},
    {'id': 63, 'name':	'104'},
    {'id': 64, 'name':	'105'},
    {'id': 65, 'name':	'106'},
    {'id': 66, 'name':	'107'},
    {'id': 67, 'name':	'108'},
    {'id': 68, 'name':	'109'},
    {'id': 69, 'name':	'110'},
    {'id': 70, 'name':	'111'},
    {'id': 71, 'name':	'112'},
    {'id': 72, 'name':	'113'},
    {'id': 73, 'name':	'114'},
    {'id': 74, 'name':	'115'},
    {'id': 75, 'name':	'120'},
    {'id': 76, 'name':	'121'},
    {'id': 77, 'name':	'122'},
    {'id': 78, 'name':	'123'},
    {'id': 79, 'name':	'cot'},
    {'id': 80, 'name':	'housing'},
    {'id': 81, 'name':	'patrol'},
    {'id': 82, 'name':	'pbbn'},
    {'id': 83, 'name':	'pbbs'},
    {'id': 84, 'name':	'pbbx'},
    {'id': 85, 'name':	'pbmn'},
    {'id': 86, 'name':	'pbms'},
    {'id': 87, 'name':	'pbqn'},
    {'id': 88, 'name':	'pbqs'},
    {'id': 89, 'name':	'pbsi'},
    {'id': 90, 'name':	'transit'}
]

# A legacy data folder must exist in the scripts folder to parse from
# TODO: Maybe name this something else?
timestamped_folders = os.listdir('legacy_data')
timestamped_folders.sort()
for folder in timestamped_folders:
    files = os.listdir('legacy_data/' + folder)
    files.sort()
    year = folder.split("_")[0]
    month = folder.split("_")[1]
    for precinct_file in files:
        if '.csv' not in precinct_file:
            continue
        precinct_short_name = precinct_file.replace('sum.csv', '')
        precinct_col = [precinct for precinct in precincts if precinct['name'] == precinct_short_name]
        precinct_id = None
        if len(precinct_col) == 1:
          precinct_id = precinct_col[0]['id']
        else:
          raise Exception('Precinct not found. Does it exist?:', precinct_short_name)

        with open('legacy_data/' + folder + '/' + precinct_file, newline='') as csvfile:
            f = csv.reader(csvfile, delimiter=',', quotechar='|')
            # Skip first row unless it's the months with flipped headers
            if folder != "2017_02_sum" and folder != "2018_08_sum":
                next(f)

            for row in f:
                # This row likely has hidden bad data
                if len(row[1]) == 0:
                    continue

                if len(row) == 4:
                    new_row = ["".join(row[:2]), row[2], row[3]]
                else:
                    new_row = row

                
                violation_count = new_row[1]

                violation_id = None
                violation_col = [violation for violation in violations if violation['name'] == new_row[0]]
                if len(violation_col) == 1:
                  violation_id = violation_col[0]['id']
                else:
                  raise Exception('Violation mapping not found for violation. Does it exist?:', new_row[0])

                stmt = "({}, {}, {}, {}, {}),".format(
                    "'"+year+"'", "'"+month+"'", "'"+violation_count+"'", precinct_id, violation_id)
                print(stmt)
                # INSERT INTO tallies (year, month, count, precinct_id, moving_violation_id) VALUES
                # (year, month, violation_count, (SELECT id from precincts WHERE short_name=precinct_short_name), (SELECT id from moving_violations WHERE name=violation_name) ),
                # (year, month, violation_count, (SELECT id from precincts WHERE short_name=precinct_short_name) );
