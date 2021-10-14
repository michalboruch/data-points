import json
import requests

from django.http import HttpResponse, JsonResponse
from django.views.decorators.csrf import csrf_exempt

from .models import DataPoint


@csrf_exempt
def data_point_list(request):
    response = HttpResponse()
    # Handle only POST method
    if not request.method == "POST":
        response.status_code = 405
        return response
    # Parse Data Points
    data = request.body
    parsed_data = []
    if data:
        try:
            data = json.loads(data)
        except Exception as e:
            print(e)
            response.status_code = 400
            return response
        else:
            parsed_data = [
                DataPoint(name=i.get("name"), timestamp=i.get("t"), value=i.get("v"))
                for i in data
            ]
    # Save Data Points
    DataPoint.objects.bulk_create(parsed_data, ignore_conflicts=True)
    response.status_code = 201
    return response


def data_point_statistics(request, name):
    # TODO: validate request data
    from_ts = request.GET.get("from")
    to_ts = request.GET.get("to")
    
    url = f"http://calc:8080/basic_stats?name={name}&from={from_ts}&to={to_ts}"

    data = requests.get(url).json()
    return JsonResponse(data)
