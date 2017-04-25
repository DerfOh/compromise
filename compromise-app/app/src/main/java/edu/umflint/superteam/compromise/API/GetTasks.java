package edu.umflint.superteam.compromise.API;

import android.os.AsyncTask;
import android.util.Log;

import org.json.JSONArray;
import org.json.JSONObject;

import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.net.HttpURLConnection;
import java.net.URL;
import java.util.ArrayList;

import edu.umflint.superteam.compromise.Classes.ExpandListChild;
import edu.umflint.superteam.compromise.Classes.ExpandListGroup;

public class GetTasks extends AsyncTask<Void, Void, ArrayList<ExpandListGroup>> {

    private final int mGroup;
    private StringBuffer response;

    public GetTasks(int group) {
        mGroup = group;
    }

    @Override
    protected ArrayList<ExpandListGroup> doInBackground(Void... params) {
        try {
            String url = "https://api.compromise.rocks/api/tasks/" + mGroup;
            URL obj = new URL(url);
            HttpURLConnection con = (HttpURLConnection) obj.openConnection();

            //add request header
            con.setRequestMethod("GET");

            int responseCode = con.getResponseCode();
            Log.i("Tasks", "Sending 'GET' request to URL : " + url);
            Log.i("Tasks", "Response Code : " + responseCode);

            BufferedReader in = new BufferedReader(
                    new InputStreamReader(con.getInputStream()));
            String inputLine;
            response = new StringBuffer();

            while ((inputLine = in.readLine()) != null) {
                response.append(inputLine);
            }
            in.close();
        } catch (Exception ex) {
            Log.e("Tasks", ex.toString());
        }

        ArrayList<ExpandListGroup> list = new ArrayList<ExpandListGroup>();
        ArrayList<ExpandListChild> list2 = new ArrayList<ExpandListChild>();
        ArrayList<ExpandListChild> list3 = new ArrayList<ExpandListChild>();

        ExpandListGroup gru1 = new ExpandListGroup();
        gru1.setName("Pending Tasks");

        list2 = new ArrayList<ExpandListChild>();

        ExpandListGroup gru2 = new ExpandListGroup();
        gru2.setName("Completed Tasks");

        JSONArray jsonArray;

        if (!response.toString().isEmpty()) {
            try {
                jsonArray = new JSONArray(response.toString());
        for (int i = 0; i < jsonArray.length(); i++) {
            JSONObject row = jsonArray.getJSONObject(i);
            if(row.getString("CompletionStatus").equals("Incomplete"))
            {
                ExpandListChild ch1_1 = new ExpandListChild();
                ch1_1.setName(row.getString("TaskName"));
                ch1_1.setTag(row.getString("TaskId"));
                ch1_1.setPoints(row.getString("PointValue"));
                ch1_1.setDescription(row.getString("TaskDescription"));
                list2.add(ch1_1);
            }
            else if (row.getString("CompletionStatus").equals("Complete")){
                ExpandListChild ch2_1 = new ExpandListChild();
                ch2_1.setName(row.getString("TaskName"));
                ch2_1.setTag(row.getString("TaskId"));
                ch2_1.setPoints(row.getString("PointValue"));
                ch2_1.setDescription(row.getString("TaskDescription"));
                list3.add(ch2_1);
            }
        }
                gru1.setItems(list2);
                gru2.setItems(list3);
                list.add(gru1);
                list.add(gru2);
            } catch (Exception ex) {
                return null;
            }

        } else {
            return null;
        }

        //print result
        Log.i("Tasks", response.toString());
        return list;
    }
}