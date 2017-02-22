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

public class GetRewards extends AsyncTask<Void, Void, ArrayList<ExpandListGroup>> {

    private final int mGroup;
    private StringBuffer response;

    public GetRewards(int group) {
        mGroup = group;
    }

    @Override
    protected ArrayList<ExpandListGroup> doInBackground(Void... params) {
        try {
            String url = "http://api.compromise.rocks/api/rewards/" + mGroup;
            URL obj = new URL(url);
            HttpURLConnection con = (HttpURLConnection) obj.openConnection();

            //add request header
            con.setRequestMethod("GET");

            int responseCode = con.getResponseCode();
            Log.i("Rewards", "Sending 'GET' request to URL : " + url);
            Log.i("Rewards", "Response Code : " + responseCode);

            BufferedReader in = new BufferedReader(
                    new InputStreamReader(con.getInputStream()));
            String inputLine;
            response = new StringBuffer();

            while ((inputLine = in.readLine()) != null) {
                response.append(inputLine);
            }
            in.close();
        } catch (Exception ex) {
            Log.e("Rewards", ex.toString());
        }

        ArrayList<ExpandListGroup> list = new ArrayList<ExpandListGroup>();
        ArrayList<ExpandListChild> list2 = new ArrayList<ExpandListChild>();
        ArrayList<ExpandListChild> list3 = new ArrayList<ExpandListChild>();

        ExpandListGroup gru1 = new ExpandListGroup();
        gru1.setName("Rewards");

        list2 = new ArrayList<ExpandListChild>();


        JSONArray jsonArray;

        if (!response.toString().isEmpty()) {
            try {
                jsonArray = new JSONArray(response.toString());
                for (int i = 0; i < jsonArray.length(); i++) {
                    JSONObject row = jsonArray.getJSONObject(i);
                    ExpandListChild ch1_1 = new ExpandListChild();
                    ch1_1.setName(row.getString("RewardName"));
                    ch1_1.setTag(row.getString("RewardId"));
                    ch1_1.setPoints(row.getString("PointCost"));
                    ch1_1.setDescription(row.getString("RewardDescription"));
                    list2.add(ch1_1);

                }
                gru1.setItems(list2);
                list.add(gru1);
            } catch (Exception ex) {
                return null;
            }

        } else {
            return null;
        }

        //print result
        Log.i("Rewards", response.toString());
        return list;
    }
}