package edu.umflint.superteam.compromise.API;

/**
 * Created by NSammut on 3/7/2017.
 */

import android.content.SharedPreferences;
import android.os.AsyncTask;
import android.preference.PreferenceManager;
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

public class GetTaskLeaders extends AsyncTask<Void, Void, Boolean> {

    private final int mGroup;
    private final String mLoggedIn;
    private StringBuffer response;

    public GetTaskLeaders(int group, String loggedIn) {
        mGroup = group;
        mLoggedIn = loggedIn;
    }

    @Override
    protected Boolean doInBackground(Void... params) {
        boolean isLeader = false;
        try {
            String url = "http://api.compromise.rocks/api/taskleaders/" + mGroup;
            URL obj = new URL(url);
            HttpURLConnection con = (HttpURLConnection) obj.openConnection();

            //add request header
            con.setRequestMethod("GET");

            int responseCode = con.getResponseCode();
            Log.i("TaskLeader", "Sending 'GET' request to URL : " + url);
            Log.i("TaskLeader", "Response Code : " + responseCode);

            BufferedReader in = new BufferedReader(
                    new InputStreamReader(con.getInputStream()));
            String inputLine;
            response = new StringBuffer();

            while ((inputLine = in.readLine()) != null) {
                response.append(inputLine);
            }
            in.close();
        } catch (Exception ex) {
            Log.e("TaskLeader", ex.toString());
        }

        JSONArray jsonArray;
        if (!response.toString().isEmpty()) {
            try {
                jsonArray = new JSONArray(response.toString());
                for (int i = 0; i < jsonArray.length(); i++) {
                    JSONObject row = jsonArray.getJSONObject(i);
                    if(row.getString("EmailAddress").equals(mLoggedIn))
                    {
                        isLeader = true;
                    }
                }
            } catch (Exception ex) {
                return isLeader;
            }

        } else {
            return isLeader;
        }

        //print result
        Log.i("TaskLeader", response.toString());
        return isLeader;
    }
}