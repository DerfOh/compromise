package edu.umflint.superteam.compromise.API;

import android.os.AsyncTask;
import android.util.Log;
import android.widget.Toast;

import org.json.JSONArray;
import org.json.JSONObject;

import java.io.BufferedReader;
import java.io.DataOutputStream;
import java.io.InputStreamReader;
import java.net.HttpURLConnection;
import java.net.URL;

import edu.umflint.superteam.compromise.Classes.ExpandListChild;
import edu.umflint.superteam.compromise.Classes.ExpandListGroup;

public class NewTask extends AsyncTask<Void, Void, String> {

    private final int mGroupId;
    private final String mTaskName;
    private final String mTaskDescription;
    private final int mPointValue;

    private StringBuffer response;

    public NewTask(int groupId, String taskName, String taskDescription, int pointValue) {
        mGroupId = groupId;
        mTaskName = taskName;
        mTaskDescription = taskDescription;
        mPointValue = pointValue;
    }

    @Override
    protected String doInBackground(Void... params) {
        try {
            String url = "https://api.compromise.rocks/api/tasks/";
            URL obj = new URL(url);
            HttpURLConnection con = (HttpURLConnection) obj.openConnection();

            //add request header
            con.setRequestMethod("POST");
            con.setRequestProperty("Content-Type", "application/x-www-form-urlencoded");

            String urlParameters = "GroupId=" + mGroupId + "&TaskName=" + mTaskName + "&TaskDescription=" + mTaskDescription + "&DateDue=2015-10-28&ApprovalStatus=Pending&CompletionStatus=Incomplete&PointValue=" + mPointValue;

            // Send post request
            con.setDoOutput(true);
            DataOutputStream wr = new DataOutputStream(con.getOutputStream());
            wr.writeBytes(urlParameters);
            wr.flush();
            wr.close();

            int responseCode = con.getResponseCode();
            Log.i("NTask", "Sending 'POST' request to URL : " + url);
            Log.i("NTask", "Response Code : " + responseCode);

            BufferedReader in = new BufferedReader(
                    new InputStreamReader(con.getInputStream()));
            String inputLine;
            response = new StringBuffer();

            while ((inputLine = in.readLine()) != null) {
                response.append(inputLine);
            }
            in.close();
        } catch (Exception ex) {
            Log.e("NTask", ex.toString());
        }


        //print result
        Log.i("NTask", response.toString());
        return response.toString();
    }
}