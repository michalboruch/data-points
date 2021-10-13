from django.db import models


class DataPoint(models.Model):
    name = models.CharField(max_length=100)
    timestamp = models.IntegerField()
    value = models.FloatField()
