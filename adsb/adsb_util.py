import os
import json
import codecs
import pandas as pd
from dataclasses import make_dataclass


class flightVector:
    def __init__(self, id):
        self.id = id
        self.lats = list()
        self.longs = list()
        self.flighttime = list()
        self.stationtime = list()
        self.latency = list()
        self.altitudes = list()


def pfs(fpath):
    files = [os.path.join(fpath, fn) for fn in os.listdir(fpath)]
    ftable = dict()
    for fname in files:
        with codecs.open(fname, 'r', encoding='utf-8', errors='ignore') as rf:
            sf = rf.readline()
        fdata = json.loads(sf)

        for fobj in fdata['acList']:
            if fobj['Id'] not in ftable:
                ftable[fobj['Id']] = flightVector(fobj['Id'])
            ftable[fobj['Id']].lats.append(fobj['Lat'])
            ftable[fobj['Id']].longs.append(fobj['Long'])
            if 'PosTime' in fobj:
                ftable[fobj['Id']].flighttime.append(fobj['PosTime'])
            else:
                ftable[fobj['Id']].flighttime.append(
                    ftable[fobj['Id']].flight[-1])
            ftable[fobj['Id']].stationtime.append(fdata['stm'])
            ftable[fobj['Id']].latency.append(
                ftable[fobj['Id']].stationtime[-1] - ftable[fobj['Id']].flighttime[-1])
            if 'Alt' in fobj:
                ftable[fobj['Id']].altitudes.append(fobj['Alt'])
            elif len(ftable[fobj['Id']].altitudes) != 0:
                ftable[fobj['Id']].altitudes.append(
                    ftable[fobj['Id']].altitudes[-1])
            else:
                ftable[fobj['Id']].altitudes.append(0)

    return ftable


def tdf(flightVectors):
    FlightPoint = make_dataclass("FlightPoint", [("Id", int), ("Lat", float), ("Long", float), (
        "Flighttime", int), ("Stationtime", int), ("Latency", int), ("Altitude", int)])
    # return pd.DataFrame([FlightPoint()])
    data = []
    for row in flightVectors.values():
        data.extend([FlightPoint(row.id, lat, long_, flighttime, stationtime, latency, altitude) for lat, long_, flighttime,
                    stationtime, latency, altitude in zip(row.lats, row.longs, row.flighttime, row.stationtime, row.latency, row.altitudes)])
    return pd.DataFrame(data)


def clean(flightDf):
    flightDf = flightDf[flightDf.Latency > 0] # latency less than 0 doesn't make sense
    return flightDf
