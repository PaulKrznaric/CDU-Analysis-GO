# CDU-Analysis-GO
This is a project that performs data analysis of a spreadsheet of patient data outputted by an EMR. 
The software is able to take in a spreadsheet of the (highly) specific format and output the billing amounts into a format that can 
then be appended onto the existing spreadsheet format. 

The software has been written in a manner such that at some point the entire process can be automated, however, it's unlikely that
the time it would take to fully automate the process would account for the time difference between the manual formatting. 

This project replaces an older version of this software written in Java when requirements were different and involved comparison: 
[cduDataExtract Java](https://github.com/PaulKrznaric/cduDataExtract)

# Run-Time Information
This project is unlikely to be relevant to you, however there is a Makefile.

In order to run this project you'll need to modify the GO Code. [Clone the repo](https://github.com/PaulKrznaric/CDU-Analysis-GO.git) and navigate in your favourite IDE to [/bin/read.go](https://github.com/PaulKrznaric/CDU-Analysis-GO.git). You'll need to modify [This line](https://github.com/PaulKrznaric/CDU-Analysis-GO.git) to wherever your input file is located. 
After that you should be able to make run the software. 
### Spreadsheet Specifications
The following will need to be the case for the spreadsheet in order for the calculation to be successful:
- Any header information's first column should be labled "MDN"
- Patient ID in column 1 
  - Due to my specific business case, the expected format is as follows (in REGEX): ^J0*\d*$
  - In human readable: 
    - Starts with the character 'J'
    - Has between zero and unlimited 0's 
      - These leading zeros are trimmed
    - Has between zero and ulimited decimal numbers (0-9)
    - Terminates
  -i.e.: "J000123" is acceptable, "J000123a" is not. 
- The primary treating doctor in column 3
- Patient time in in column 4*
- The secondary treating docter in column 8
- The time the patient left the CDU unit in column 9* 
- The time the patient was admitted to the hopsital in column 11*
  - If the patient was not admitted, please leave this line blank
-The entire length of the row is 14
  - This is to catch any issues for _my personal business use_ at runtime
  
*This is expected in the YYYY-MM-DD H:MM:SS AM/PM format

### Example:

| MDN     | ??? | ??? | Primary Doctor | Patient Time In      | ??? | ??? | ??? | Secondary Doctor | Depart CDU Time      | ??? | Admitted Time        | ??? | ??? |
|---------|-----|-----|----------------|----------------------|-----|-----|-----|------------------|----------------------|-----|----------------------|-----|-----|
| J000123 |     |     | Dr. John Doe   | 2022-8-28 2:10:00 PM |     |     |     | Dr. Jane Doe     | 2022-8-28 3:00:00 PM |     | 2022-8-28 2:30:00 PM |     |     |

Where ??? represents irrelevant data for this import. 
