# Trade module for buy/sell operations

    This module contains following classes (structs):
- TradePool - this is a struct that contain all the buy&sell operations of specific market, it should be created once for each new market, and then called when some new operations of buy/sell are added/cancelled on some market
- Buy - this is a buy operation, this operation refers to buying a specific market asset, and selling main asset
- Sell - this is a sell operation, this operation refers to selling a specific market asset, and buying main asset
- Output - this is an output operation made to transfer information about the outputs that should be operated to some users after some operations being executed

# Creation of trade pool

    Single example of trade pool should be created once for each market. Each
trade pool has methods to add buy/sell trade, and those methods return oputputs,
that should be operated on market level(that is made by such a way, because market should not have any depecdencies to user module).

# Adding/Cancelling trades

    It is possible to add buy/sell trades to each trade pool. Each added trade 
will firstly be operated on existing trades (if there are some), in case if `match` operation didn't cause any result, the trade is gonna be putted to market pool.
Each added trade can be cancelled, if trade is cancelled all current trade's `offer` will be returned to trade's creator.

# Trades matching

    When trade is added to some TradePool, firstly it's being matched with
other trades. It is working in such a way, that if buy matched to sell/ sell matched to buy are both profitable for buyer and seller, then trades are being operated. If matching is having overage on buy trade offer, it goes to seller. If matching is having overage on sell trade offer, it goes to buyer. If there is overage on both buyer and seller sides, then buyers overage goes to seller, while sellers overage goes to buyer.