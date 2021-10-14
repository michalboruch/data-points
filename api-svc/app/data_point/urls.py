from django.urls import path

from . import views


urlpatterns = [
    path("", views.data_point_list),
    path("<str:name>", views.data_point_statistics),
]
