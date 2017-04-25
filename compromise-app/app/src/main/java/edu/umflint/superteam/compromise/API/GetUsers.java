package edu.umflint.superteam.compromise.API;

import android.os.AsyncTask;
import android.support.design.widget.NavigationView;
import android.util.Log;
import android.view.Menu;

import org.json.JSONArray;
import org.json.JSONObject;

import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.net.HttpURLConnection;
import java.net.URL;
import java.util.ArrayList;

import edu.umflint.superteam.compromise.R;

public class GetUsers extends AsyncTask<Void, Void, ArrayList> {

    private final int mGroup;
    private StringBuffer response;

    public GetUsers(int group) {
        mGroup = group;
    }

    @Override
    protected ArrayList doInBackground(Void... params) {
        // TODO: attempt authentication against a network service.

        try {
            String url = "https://api.compromise.rocks/api/users/" + mGroup;
            URL obj = new URL(url);
            HttpURLConnection con = (HttpURLConnection) obj.openConnection();

            int responseCode = con.getResponseCode();
            Log.i("Users", "Sending 'GET' request to URL : " + url);
            Log.i("Users", "Response Code : " + responseCode);

            BufferedReader in = new BufferedReader(
                    new InputStreamReader(con.getInputStream()));
            String inputLine;
            response = new StringBuffer();

            while ((inputLine = in.readLine()) != null) {
                response.append(inputLine);
            }
            in.close();
        } catch (Exception ex) {
            Log.e("Users", ex.toString());
            return null;
        }

        //print result
        Log.i("Users", response.toString());
        if (!response.toString().isEmpty()) {
            try {
                JSONArray jsonObj = new JSONArray(response.toString());
                ArrayList names = new ArrayList();

                String firstName;
                String lastName;
                int totalPoints;
                String nickName;

                try {
                    if (jsonObj != null) {

                        for (int i = 0; i < jsonObj.length(); i++) {
                            JSONObject row = jsonObj.getJSONObject(i);
                            firstName = row.getString("FirstName");
                            lastName = row.getString("LastName");
                            nickName = row.getString("NickName");
                            totalPoints = row.getInt("TotalPoints");
                            names.add(firstName + " " + lastName + "\t|\t Points: " + totalPoints);
                        }
                        return names;

                    } else {

                    }
                } catch (Exception ex) {

                }

            } catch (Exception ex) {
                return null;
            }

        } else {
            return null;
        }
        return null;
    }
}