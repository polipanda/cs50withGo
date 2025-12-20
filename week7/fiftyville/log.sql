-- Keep a log of any SQL queries you execute as you solve the mystery.
-- All you know is that the theft took place on July 28, 2024
--      and that it took place on Humphrey Street.
SELECT *
FROM crime_scene_reports
WHERE street = 'Humphrey Street'
AND year = 2024
AND month = 7
AND day = 28;
-- id is 295
-- Theft of the CS50 duck took place at 10:15am at the Humphrey Street bakery. Interviews were conducted today with
-- three witnesses who were present at the time – each of their interview transcripts mentions the bakery.

SELECT *
FROM interviews
WHERE year = 2024 AND month = 7 AND day = 28;
-- 7 transcripts, 3 mention the bakery:
--  Sometime within ten minutes of the theft, I saw the thief get into a car in the bakery parking lot and
--  drive away. If you have security footage from the bakery parking lot, you might want to look for cars that
--  left the parking lot in that time frame.

--  I don't know the thief's name, but it was someone I recognized. Earlier this morning, before I arrived at Emma's
--  bakery, I was walking by the ATM on Leggett Street and saw the thief there withdrawing some money.

--  As the thief was leaving the bakery, they called someone who talked to them for less than a minute. In the
--  call, I heard the thief say that they were planning to take the earliest flight out of Fiftyville tomorrow.
--  The thief then asked the person on the other end of the phone to purchase the flight ticket.

-- pulling peoples phone numbers for anyone making a withdrawal at Leggett Street the morning of the theft
-- and license plate matches the bank security logs around the time of the theft
--
SELECT people.phone_number FROM bank_accounts
JOIN people ON bank_accounts.person_id = people.id
WHERE account_number IN
      (SELECT account_number FROM atm_transactions
       WHERE year = 2024 AND month = 7 AND day = 28 AND atm_location = 'Leggett Street' and transaction_type = 'withdraw')
AND license_plate IN
    (SELECT license_plate FROM bakery_security_logs
    WHERE year = 2024 AND month = 7 AND day = 28 AND hour = 10 AND minute > 15 AND minute < 30 AND activity = 'exit')
;
-- possible phone numbers: (367) 555-5533, (770) 555-1861, (829) 555-5269, (389) 555-5198

-- checking passports for people who made a call on the day of the theft, from the
SELECT passport_number
FROM people
WHERE phone_number IN (
    SELECT caller FROM phone_calls
    WHERE year = 2024 AND month = 7 AND day = 28
    AND caller IN ('(367) 555-5533', '(770) 555-1861', '(829) 555-5269', '(389) 555-5198)')
    GROUP BY caller
    );
-- possible numbers are: 3592750733, 5773159633

-- if we get one person from checking the earliest flight out, then we know who did it
SELECT name FROM people WHERE passport_number IN (
    SELECT passengers.passport_number
    FROM passengers
    WHERE flight_id = (
        SELECT id FROM flights
        WHERE origin_airport_id = (SELECT id FROM airports WHERE city = 'Fiftyville')
          AND year = 2024 AND month = 7 AND day = 29
        ORDER BY hour
        LIMIT 1
    )
      AND passport_number IN (3592750733, 5773159633)
    );
-- Bruce did it

-- Flight he escaped to is New York City
SELECT city
FROM airports
WHERE id = (SELECT destination_airport_id from flights where id = 36);

-- Accomplice is someone Bruce called for less than a minute
SELECT name
FROM people
WHERE phone_number IN (
    SELECT receiver FROM phone_calls
    WHERE caller = '(367) 555-5533'
    AND year = 2024 AND month = 7 AND day = 28
    AND duration < 60
);

-- Robin
