# cGem

A simple app which allows you to quickly buy and sell cryto on the Gemini Exchange.

## Clone and Install

```bash
git clone https://github.com/james-daniels/cgem.git

cd cgem
```

```bash
$ ./cgem -h
Use cGem to quickly buy and sell cryto on the Gemini Exchange

Usage:
  cGem [command]

Available Commands:
  buy         Buy side places order to buy crypto
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  price       Price of trading pair
  sell        Sell side places order to sell crypto

Flags:
  -h, --help   help for cGem

Use "cGem [command] --help" for more information about a command.
```

## Easy process to buy

```bash
$ ./cgem buy -h
Buy will fill part of the order it can immediately, then cancel any remaining amount.

Usage:
  cGem buy [flags]

Flags:
  -a, --amount int      amount to buy
  -h, --help            help for buy
  -o, --offset int      amount to add to price
  -s, --symbol string   symbol of the trading pair
```

The default value for --offset is zero (0) and is optional. However, in order to fill an order, the offset value to increase the purchase price aggressive enough to fill the order.

```bash
$ ./cgem buy -s ltcusd -a 1 -o 5

OrderID:                1772530063
ID:                     1772530063
Symbol:                 ltcusd
Exchange:               gemini
AvgExecutionPrice:      63.95
Side:                   buy
Type:                   exchange limit
Timestamp:              1652324695
Timestampms:            1652324695332
IsLive:                 false
IsCancelled:            false
IsHidden:               false
WasForced:              false
ExecutedAmount:         1
Options:                [immediate-or-cancel]
StopPrice:
Price:                  68.82
OriginalAmount:         1
```

## Easy process to sell

```Bash
$ ./cgem sell -h
Sell will fill part of the order it can immediately, then cancel any remaining amount.

Usage:
  cGem sell [flags]

Flags:
  -a, --amount int      amount to sell
  -h, --help            help for sell
  -o, --offset int      amount to add
  -s, --symbol string   symbol of the trading pair
```

Reverse the buy process to sell.  Offset needs to be negative value which is added to the sell price.

```bash
$ ./cgem sell -s ltcusd -a 1 -o -5

OrderID:                1772532056
ID:                     1772532056
Symbol:                 ltcusd
Exchange:               gemini
AvgExecutionPrice:      63.82
Side:                   sell
Type:                   exchange limit
Timestamp:              1652324782
Timestampms:            1652324782601
IsLive:                 false
IsCancelled:            false
IsHidden:               false
WasForced:              false
ExecutedAmount:         1
Options:                [immediate-or-cancel]
StopPrice:
Price:                  58.88
OriginalAmount:         1
```

## Get the price of your favorite crypto

```bash,
$ ./cgem price -s ltcusd

LTCUSD: 63.97
```

## Config file

The config.ini file needs be located in the same directory as the executable.

```ini
#Possible values: sandbox and production
environment = sandbox

#Optional: Present output in human readable format
#Only available for single run jobs
#pretty = true

[credentials]
#API key and secret
apikey = account-XXXXXXXXXXXXXXXXXXXX
apisecret = XXXXXXXXXXXXXXXXXXXX

[recurrence]
#Optional: Only for recurring jobs
#repeat = false

#Dependent on repeat = true
#Number of hours between runs
#frequency = 0

[orders]
#Default value is 0
#The API does not support market orders because it does not provide price protection.
#Offset agressively coupled with the curret price increases or decreases the limit price.
#This will achieve the same result as a market order.
#offset = 0

[logging]
#Optional: path to log file location
; logfile = "cgem.log"
```
