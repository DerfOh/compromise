package edu.umflint.superteam.compromise.API;

import android.os.AsyncTask;
import android.util.Log;

import org.json.JSONArray;
import org.json.JSONObject;

import java.io.BufferedReader;
import java.io.DataOutputStream;
import java.io.InputStreamReader;
import java.net.HttpURLConnection;
import java.net.URL;

import edu.umflint.superteam.compromise.Classes.ExpandListChild;
import edu.umflint.superteam.compromise.Classes.ExpandListGroup;

public class GetPassword extends AsyncTask<Void, Void, String> {

    private final String mEmail;
    private StringBuffer response;

    public GetPassword(String email) {
        mEmail = email;
    }

    @Override
    protected String doInBackground(Void... params) {
        try {
            String url = "https://api.compromise.rocks/api/retrievepassword/";
            URL obj = new URL(url);
            HttpURLConnection con = (HttpURLConnection) obj.openConnection();

            //add request header
            con.setRequestMethod("POST");
            con.setRequestProperty("Content-Type", "application/x-www-form-urlencoded");

            String urlParameters = "EmailAddress=" + mEmail;

            // Send post request
            con.setDoOutput(true);
            DataOutputStream wr = new DataOutputStream(con.getOutputStream());
            wr.writeBytes(urlParameters);
            wr.flush();
            wr.close();

            int responseCode = con.getResponseCode();
            Log.i("Password", "Sending 'POST' request to URL : " + url);
            Log.i("Password", "Response Code : " + responseCode);

            BufferedReader in = new BufferedReader(
                    new InputStreamReader(con.getInputStream()));
            String inputLine;
            response = new StringBuffer();

            while ((inputLine = in.readLine()) != null) {
                response.append(inputLine);
            }
            in.close();
        } catch (Exception ex) {
            Log.e("Password", ex.toString());
        }


        //print result
        Log.i("Tasks", response.toString());
        return response.toString();
    }
}