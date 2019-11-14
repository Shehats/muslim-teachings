from flask import Flask
from flask_graphql import GraphQLView
from models.base_model import db_session
from app_queries import all_schema
from quran_csv_to_db import QuranCSVToDB

app = Flask(__name__)
app.debug = True

app.add_url_rule(
    '/api/quran',
    view_func=GraphQLView.as_view(
        'quran',
        schema=all_schema,
        graphiql=True
    )
)

@app.teardown_appcontext
def shutdown_session(exception=None):
    db_session.remove()

if __name__ == '__main__':
    path = "{}/resources/Quran.csv".format(os.path.dirname(os.path.abspath(__file__)))
    audio_path = "{}/resources/audio".format(os.path.dirname(os.path.abspath(__file__)))
    q = QuranCSVToDB(path)
    # q.walk_audio_files_non_ordered_numeric(audio_path, "alafasy")
    print(q.load_recitations_files_from_file(audio_path, "alafasy"))
    app.run()